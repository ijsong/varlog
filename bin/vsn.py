#!/usr/bin/python
# -*- coding: utf-8 -*-

import grp
import json
import os
import pwd
import shutil
import socket
import subprocess
import sys
import time
from random import randint

cwd = os.path.dirname(os.path.realpath(__file__))  # noqa
binpath = os.path.join(cwd, "..", "bin")  # noqa
pylib = os.path.join(cwd, "..", "pylib")  # noqa
if os.path.isdir(pylib):  # noqa
    sys.path.insert(0, pylib)  # noqa

from varlog.killer import Killer  # noqa
from varlog import procutil  # noqa
from varlog import limits  # noqa
from varlog.logger import get_logger  # noqa

logger = get_logger("vsn")

RETRY_INTERVAL_SEC = 3
ADD_STORAGE_NODE_INTERVAL_SEC = 1
DEFAULT_CLUSTER_ID = "1"
DEFAULT_RPC_PORT = "9091"
DEFAULT_VSN_HOME = "/home/deploy/varlog-sn"
TEST_STORAGE = "/home/deploy/storage"
LOCAL_ADDRESS = socket.gethostbyname(socket.gethostname())


def get_advertise_addr():
    host = os.getenv("HOST_IP", LOCAL_ADDRESS)
    port = os.getenv("RPC_PORT", DEFAULT_RPC_PORT)
    return f"{host}:{port}"


def get_vms_addr():
    addr = os.getenv("VMS_ADDRESS")
    if addr is None:
        raise Exception("no vms address")
    return addr


def get_storage_node_id():
    try:
        out = subprocess.check_output(
            [f"{binpath}/vmc", "meta", "sn",
             "--vms-address=%s" % get_vms_addr()])
        meta = json.loads(out)
        my_addr = get_advertise_addr()
        for snid, addr in meta["storagenodes"].items():
            if addr == my_addr:
                return snid, True
        return randint(1, (2 ^ 31) - 1), False
    except Exception as e:
        logger.exception("could not get StorageNodeID")
        raise e


def get_volume(truncate=False):
    home = os.getenv("VSN_HOME", DEFAULT_VSN_HOME)
    datapath = f"{home}/data"

    if truncate:
        try:
            shutil.rmtree(datapath)
        except FileNotFoundError:
            pass

    os.makedirs(datapath, exist_ok=True)
    return datapath


def add_storage_node(addr):
    try:
        subprocess.check_output(
            [f"{binpath}/vmc", "add", "sn", f"--storage-node-address={addr}",
             "--vms-address=%s" % get_vms_addr()])
    except subprocess.CalledProcessError as e:
        logger.exception("could not add storagenode")
        raise e


def main():
    limits.set_limits()
    cluster_id = os.getenv("CLUSTER_ID", DEFAULT_CLUSTER_ID)
    listen_port = os.getenv("RPC_PORT", DEFAULT_RPC_PORT)
    snid, exist = get_storage_node_id()
    volume = get_volume(truncate=not exist)
    listen_addr = f"0.0.0.0:{listen_port}"
    advertise_addr = get_advertise_addr()

    storage_node = [
        f"{binpath}/vsn",
        "start",
        f"--cluster-id={cluster_id}",
        f"--storage-node-id={snid}",
        f"--listen-address={listen_addr}",
        f"--volumes={volume}",
        f"--advertise-address={advertise_addr}"
    ]

    killer = Killer()
    while not killer.kill_now:
        if procutil.check_liveness("vsn"):
            time.sleep(RETRY_INTERVAL_SEC)
            continue
        try:
            procutil.kill("vsn")
            logger.info(f"running storage node: {storage_node}")
            subprocess.Popen(storage_node)
            time.sleep(ADD_STORAGE_NODE_INTERVAL_SEC)
            if not exist:
                logger.info("adding storage node to cluster")
                add_storage_node(advertise_addr)
        except (OSError, ValueError, subprocess.SubprocessError):
            logger.exception("could not run storage node")
        time.sleep(RETRY_INTERVAL_SEC)
    procutil.stop("vsn")


if __name__ == "__main__":
    main()
