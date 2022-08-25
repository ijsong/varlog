# Changelog

## [0.0.1](https://github.com/ijsong/varlog/compare/v0.0.3...v0.0.1) (2022-08-25)


### Features

* **admin:** check data path to run storagenode ([#693](https://github.com/ijsong/varlog/issues/693)) ([9928c19](https://github.com/ijsong/varlog/commit/9928c19538c20b8f96248086cd928fff66cf02b0))
* **storagenode:** check required data directories ([#694](https://github.com/ijsong/varlog/issues/694)) ([117be34](https://github.com/ijsong/varlog/commit/117be34bcbc4f7da69dcc54468b55eef7e8ca4aa))


### Bug Fixes

* add cancel to context.WithTimeout to avoid context leak ([#307](https://github.com/ijsong/varlog/issues/307)) ([c756aa9](https://github.com/ijsong/varlog/commit/c756aa9c176b72246c81a2993fbd8a7bede5cd2d))
* **admin:** add handler timeout for failed sn ([a2f31d7](https://github.com/ijsong/varlog/commit/a2f31d7b7b43a8522dd513ac824d040a6f515217)), closes [#29](https://github.com/ijsong/varlog/issues/29)
* **mr:** let newbie logstream know cur version ([cd12789](https://github.com/ijsong/varlog/commit/cd12789f91fe1e4d17b14a8636535612a3fc793b))
* race condition in SNManager ([#302](https://github.com/ijsong/varlog/issues/302)) ([f58526d](https://github.com/ijsong/varlog/commit/f58526d0e7a05204138668968c659b6cbf7f0832))
* remove mutex in storage node manager of admin ([77ed718](https://github.com/ijsong/varlog/commit/77ed7188883b488c48c89812548e9c6f5c889649)), closes [#30](https://github.com/ijsong/varlog/issues/30)
* TestAdmin_GetStorageNode_FailedStorageNode ([#13](https://github.com/ijsong/varlog/issues/13)) ([5c8a3c2](https://github.com/ijsong/varlog/commit/5c8a3c234032e3bf647d2a5d10c9916c215a6d9b))


### Miscellaneous Chores

* release 0.0.1 ([#21](https://github.com/ijsong/varlog/issues/21)) ([6aad0d8](https://github.com/ijsong/varlog/commit/6aad0d80d7f3c00092d44bbcdad7730e6e956870))

## [0.0.3](https://github.com/kakao/varlog/compare/v0.0.2...v0.0.3) (2022-08-25)


### Bug Fixes

* **mr:** let newbie logstream know cur version ([cd12789](https://github.com/kakao/varlog/commit/cd12789f91fe1e4d17b14a8636535612a3fc793b))

## [0.0.2](https://github.com/kakao/varlog/compare/v0.0.1...v0.0.2) (2022-08-17)


### Bug Fixes

* **admin:** add handler timeout for failed sn ([a2f31d7](https://github.com/kakao/varlog/commit/a2f31d7b7b43a8522dd513ac824d040a6f515217)), closes [#29](https://github.com/kakao/varlog/issues/29)
* remove mutex in storage node manager of admin ([77ed718](https://github.com/kakao/varlog/commit/77ed7188883b488c48c89812548e9c6f5c889649)), closes [#30](https://github.com/kakao/varlog/issues/30)

## 0.0.1 (2022-08-14)


### Bug Fixes

* TestAdmin_GetStorageNode_FailedStorageNode ([#13](https://github.com/kakao/varlog/issues/13)) ([5c8a3c2](https://github.com/kakao/varlog/commit/5c8a3c234032e3bf647d2a5d10c9916c215a6d9b))


### Miscellaneous Chores

* release 0.0.1 ([#21](https://github.com/kakao/varlog/issues/21)) ([6aad0d8](https://github.com/kakao/varlog/commit/6aad0d80d7f3c00092d44bbcdad7730e6e956870))
