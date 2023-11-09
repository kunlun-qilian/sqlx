# Change Log

All notable changes to this project will be documented in this file.
See [Conventional Commits](https://conventionalcommits.org) for commit guidelines.



# 0.0.1

### Bug Fixes

* **fix** 🐛 解决分区被当做索引处理问题 ([12d2768](https://github.com/kunlun-qilian/sqlx/commit/12d2768c34c134066b97fbb6e20c784d71065009))
* **fix** 🐛 add alias ([cb4d72e](https://github.com/kunlun-qilian/sqlx/commit/cb4d72ef0e91031d2ca918446d24c42acd354731))
* **fix** key index expr update ([9328961](https://github.com/kunlun-qilian/sqlx/commit/932896149ddf8deb7c797b0e38e7d276547aaa46))
* **fix** tests ([cfeef6a](https://github.com/kunlun-qilian/sqlx/commit/cfeef6ac3b9f03bef22328423676b1b783503a21))
* **fix** mysql support database name with reserved words ([0adc6b2](https://github.com/kunlun-qilian/sqlx/commit/0adc6b2dc8e92c43202e5a5ea56d1421d7d542b3))
* **fix** typo ([79d97c4](https://github.com/kunlun-qilian/sqlx/commit/79d97c41e1317760b5e312ff6c18ed18ab5115a2))
* **fix** drop index ([f0a09b1](https://github.com/kunlun-qilian/sqlx/commit/f0a09b165417c85514db25054349cb856fed77f5))
* **fix(builder):** EachStructField fix ([1b71e9b](https://github.com/kunlun-qilian/sqlx/commit/1b71e9bdc7c9f37643b6322f357e635d907c3066))
* **fix(builder):** loc pollution fix ([ad35502](https://github.com/kunlun-qilian/sqlx/commit/ad355029c3a8816bdee796c5c382d0efaa69c0f5))
* **fix(builder):** reflectx should use x/reflect ([7d69bea](https://github.com/kunlun-qilian/sqlx/commit/7d69bea5f35b39862a0cc1fde4d980ee7476d148))
* **fix(builder):** TableFiels should store loc of model value ([41843c2](https://github.com/kunlun-qilian/sqlx/commit/41843c2b8d86f4354355efc1ecc68e7e93f83bb8))
* **fix** revert expr from context ([8c9f485](https://github.com/kunlun-qilian/sqlx/commit/8c9f48583a347156d1d7d5b5f713f3cceca4a229))
* **fix(connectors/postgresql):** should have blank before index def ([b2601de](https://github.com/kunlun-qilian/sqlx/commit/b2601de9620b73151dd103acb38107ab1b1678cf))
* **fix** nil cause ([8bd5e94](https://github.com/kunlun-qilian/sqlx/commit/8bd5e948090e546e4cc5dd76b71cb971df2bfe18))
* **fix** error ([ab4ddaf](https://github.com/kunlun-qilian/sqlx/commit/ab4ddaf98b014be877743932efb0c1054d80e960))
* **fix** nil 'cause' ([8a71dea](https://github.com/kunlun-qilian/sqlx/commit/8a71deaf4a77f69260d0206c7547f15c15e366c9))
* **fix** should ignore deprecated field value ([9d393eb](https://github.com/kunlun-qilian/sqlx/commit/9d393eb782d773ae2c7bce4abdfde0401770a4c1))
* **fix** should ignore deprecated field value ([acdc920](https://github.com/kunlun-qilian/sqlx/commit/acdc9205aa14e11b0ef06f24226232d8d9fc2a52))
* **fix** should UnwrapAll before db err check. ([6503ee0](https://github.com/kunlun-qilian/sqlx/commit/6503ee04c06eb296fcbc9c2e8b0abe5ad8ea263a))
* **fix** driver connect issue when ctx pass may be cancel ([da0dbe4](https://github.com/kunlun-qilian/sqlx/commit/da0dbe4cbbdea220082fa6a6fa64a7b04edf7c22))
* **fix(pg):** avoid gen invalid cmd ([42e911f](https://github.com/kunlun-qilian/sqlx/commit/42e911fcdfcba0f5501fa5878e0b2e9654f23fcb))
* **fix(migration):** create table should dry run ([93c0b30](https://github.com/kunlun-qilian/sqlx/commit/93c0b304938fa4ef8971412891d77964f6c80f5a))
* **fix(pg):** default number should be with quote and dataType ([4c0575b](https://github.com/kunlun-qilian/sqlx/commit/4c0575bf3cc9a1b03ee4715dd0fc9ffabebeb4ed))
* **fix** pg comma fix ([9b52cbe](https://github.com/kunlun-qilian/sqlx/commit/9b52cbebed06597a72fe5feff207aacc1b803583))
* **fix(migration):** log prev default value ([1cce629](https://github.com/kunlun-qilian/sqlx/commit/1cce629042e1877fcdb9c3f8304b00af547444cc))
* **fix** enhance GetColumnName ([d7074af](https://github.com/kunlun-qilian/sqlx/commit/d7074af444c88c61f547a37636d294b18b94c0ee))
* **fix(builder):** fieldValue fix ([ab91aec](https://github.com/kunlun-qilian/sqlx/commit/ab91aecbb684be35ec597eae670759039ac70c6b))
* **fix(builder):** only struct could be visited ([ef21879](https://github.com/kunlun-qilian/sqlx/commit/ef218797c8245462060a6fd45b4bf6cc84e81cd2))
* **fix** table diff ([fa57361](https://github.com/kunlun-qilian/sqlx/commit/fa573616e8a3c810f7f41cecdec48108c2cabd78))
* **fix(datatypes):** timestamp and datetime string should '' when zero ([d4ba5bd](https://github.com/kunlun-qilian/sqlx/commit/d4ba5bdceaa75e8190d7342b10a29a43014f74c6))
* **fix(postgresqlconnector):** fix ([66a5fe1](https://github.com/kunlun-qilian/sqlx/commit/66a5fe18a83bba87b2a8e9518a827f96d56a11e4))
* **fix** struct scan should skip field with type interface ([6fd72d1](https://github.com/kunlun-qilian/sqlx/commit/6fd72d147b57220db9f015099fb0cb32b22d5bd9))
* **fix(postgresqlconnector):** drop index with IF EXISTS ([7acdb4f](https://github.com/kunlun-qilian/sqlx/commit/7acdb4f89e21721f39c6fab9657744132a93ce11))
* **fix(builder):** table.Expr could use # for self holder ([d64b2b4](https://github.com/kunlun-qilian/sqlx/commit/d64b2b470e6b087e31948df8a83e7f342197b8dc))
* **fix(generator):** when with method disable should not gen ConditionByStruct ([369201a](https://github.com/kunlun-qilian/sqlx/commit/369201ae4118a5a143083a99005ce88202b9c052))
* **fix(builder):** condition builder fix ([3b75ddd](https://github.com/kunlun-qilian/sqlx/commit/3b75ddd12c9d2fb5b154675b986bd7f6bc4d2d33))
* **fix** sqlErrTypeNotFound should not for scanIterator ([e795ed9](https://github.com/kunlun-qilian/sqlx/commit/e795ed99c0fd68e66a7f09f6de62ab2989708ff1))
* **fix(generator):** Create without lastInsertId ([0e2593a](https://github.com/kunlun-qilian/sqlx/commit/0e2593a4e7e617a1aa3b4fef312c69586bab6768))
* **fix** drop table if exists ([5ee873c](https://github.com/kunlun-qilian/sqlx/commit/5ee873cf2ac56b4b1223ff729d831bea29a14898))
* **fix(builder):** force lower colume name when created ([a93b6ce](https://github.com/kunlun-qilian/sqlx/commit/a93b6ce3c8914f144891b80ff6f9afa19b59a4c2))
* **fix** scan colume match ([de64fb9](https://github.com/kunlun-qilian/sqlx/commit/de64fb9f09edfc4b46b2113e5d07f6735131dd76))
* **fix(builder):** nil handle ([685c42b](https://github.com/kunlun-qilian/sqlx/commit/685c42b268bae5422f9dd387d6f92dba081d5cd0))
* **fix(generator):** updatedAt ([72063b5](https://github.com/kunlun-qilian/sqlx/commit/72063b5fea9ab10c2cac2daa38759070036e1ab4))
* **fix(generator):** batch methods with correctly types ([340da07](https://github.com/kunlun-qilian/sqlx/commit/340da0799423060af7e13976d62c0f75f440c545))
* **fix(loggerdriver):** cost should be with unit ([cda4c7b](https://github.com/kunlun-qilian/sqlx/commit/cda4c7b0dcaea966368d32448177f9d82a644ee2))
* **fix** builder.Columns should be ptr from functions ([44252d4](https://github.com/kunlun-qilian/sqlx/commit/44252d4c1ad349d3177fb9f6517d384916d4d9a3))
* **fix** table may not be used in Count and List ([5b2d854](https://github.com/kunlun-qilian/sqlx/commit/5b2d854244143ab5f06e3efba1e1c6419f871b99))
* **fix(generator):** fix issue when no auto increment ([6ff5323](https://github.com/kunlun-qilian/sqlx/commit/6ff532335279a7df21e21440bab6c3544a6a7fe0))


### Features

* **feat** 🎸 本地化,与go-courier断联 ([e733d22](https://github.com/kunlun-qilian/sqlx/commit/e733d2200c63a7924d5b07e71186b4d7473a19d5))
* **feat** 🎸 import third party json package ([33b4069](https://github.com/kunlun-qilian/sqlx/commit/33b40690d56eb2ecca17df446d4af9f2ae02c86e))
* **feat** 🎸 修改pg indexes 查询方式 ([c4d4d95](https://github.com/kunlun-qilian/sqlx/commit/c4d4d957873a4268fd6bad79669be5d495d36d1e))
* **feat** 🎸 优化 sql 样式 ([7550914](https://github.com/kunlun-qilian/sqlx/commit/75509148118c0c9bbbd52be12bb8b4701f0ef747))
* **feat** 🎸 generate support partition ([c914a65](https://github.com/kunlun-qilian/sqlx/commit/c914a65d11559119a87d682c67f32ee85f0e755b))
* **feat** 🎸 实现分区表 migrate ([9a54e38](https://github.com/kunlun-qilian/sqlx/commit/9a54e3862c05cec865203283ecf42e5f1ea7ca1b))
* **feat** bool add OpenAPISchemaType ([a6811e7](https://github.com/kunlun-qilian/sqlx/commit/a6811e767a3474a9289a21570f944b911defdff8))
* **feat** custom index def ([e5a928c](https://github.com/kunlun-qilian/sqlx/commit/e5a928c666d0234ecc5f1868afc7885e6477c94e))
* **feat** bump to golang 1.17 and deps updates ([17fcb58](https://github.com/kunlun-qilian/sqlx/commit/17fcb5811c03ca3ab2713030581586c5669115f6))
* **feat** drop logrus ([b4a9e7b](https://github.com/kunlun-qilian/sqlx/commit/b4a9e7ba17de52967d6d29064fa97cfee05a3383))
* **feat** migration enhancement ([1fa4a9e](https://github.com/kunlun-qilian/sqlx/commit/1fa4a9e92cf79f359f3820763eeaad36b8666ea1))
* **feat** alias tag for table join rename ([fc2ed04](https://github.com/kunlun-qilian/sqlx/commit/fc2ed0417fadb395d937b684bf31b1425051b924))
* **feat** support scan joined values ([f160fe4](https://github.com/kunlun-qilian/sqlx/commit/f160fe4750d488e50566c475c88069e2c1de652b))
* **feat(builder):** support insert into select from ([85c3a84](https://github.com/kunlun-qilian/sqlx/commit/85c3a849bf4f204d30bd31b5db1079f374706eea))
* **feat(generator):** type model iterator ([40f6ab6](https://github.com/kunlun-qilian/sqlx/commit/40f6ab6c54d0c80b2fcda911a444ecfda558c87c))
* **feat(builder):** MustFields ([7e9d058](https://github.com/kunlun-qilian/sqlx/commit/7e9d0582f1fbca2ab7bad4288cb5a42e6fe5adf0))
* **feat** interface ScanIterator for scaning on flying ([1c91819](https://github.com/kunlun-qilian/sqlx/commit/1c91819b727cba8abb234df10f6cb350579e2aa2))
* **feat(generator):** switch for gen TableName or not ([c13778d](https://github.com/kunlun-qilian/sqlx/commit/c13778dce8a92d984b2eba3d10f06cdd6ced8fb8))
* **feat** feature with enum and table summary and desc ([3f2f2d9](https://github.com/kunlun-qilian/sqlx/commit/3f2f2d985eea1f695128c466d2f6325a26582e04))
* **feat(generator):** tag generator should handle DataType() string as sql data type ([b3685ed](https://github.com/kunlun-qilian/sqlx/commit/b3685ede90d8dcab19d4a77cb123a9adcca2f1a7))
* **feat(builder):** added spatial index supported and with added interface ValuerExpr for value wrapper ([fe9dfdc](https://github.com/kunlun-qilian/sqlx/commit/fe9dfdcf34b00e8d49ac6f45aa6f16035a58b63d))



# [2.23.8](https://github.com/go-courier/sqlx/compare/v2.23.7...v2.23.8)



# [2.23.7](https://github.com/go-courier/sqlx/compare/v2.23.6...v2.23.7)

### Bug Fixes

* **fix(builder):** EachStructField fix ([1b71e9b](https://github.com/go-courier/sqlx/commit/1b71e9bdc7c9f37643b6322f357e635d907c3066))



# [2.23.6](https://github.com/go-courier/sqlx/compare/v2.23.5...v2.23.6)

### Bug Fixes

* **fix(builder):** loc pollution fix ([ad35502](https://github.com/go-courier/sqlx/commit/ad355029c3a8816bdee796c5c382d0efaa69c0f5))



# [2.23.5](https://github.com/go-courier/sqlx/compare/v2.23.4...v2.23.5)

### Bug Fixes

* **fix(builder):** reflectx should use x/reflect ([7d69bea](https://github.com/go-courier/sqlx/commit/7d69bea5f35b39862a0cc1fde4d980ee7476d148))



# [2.23.4](https://github.com/go-courier/sqlx/compare/v2.23.3...v2.23.4)

### Bug Fixes

* **fix(builder):** TableFiels should store loc of model value ([41843c2](https://github.com/go-courier/sqlx/commit/41843c2b8d86f4354355efc1ecc68e7e93f83bb8))



# [2.23.3](https://github.com/go-courier/sqlx/compare/v2.23.2...v2.23.3)

### Bug Fixes

* **fix** revert expr from context ([8c9f485](https://github.com/go-courier/sqlx/commit/8c9f48583a347156d1d7d5b5f713f3cceca4a229))



# [2.23.2](https://github.com/go-courier/sqlx/compare/v2.23.1...v2.23.2)



# [2.23.1](https://github.com/go-courier/sqlx/compare/v2.23.0...v2.23.1)

### Bug Fixes

* **fix(connectors/postgresql):** should have blank before index def ([b2601de](https://github.com/go-courier/sqlx/commit/b2601de9620b73151dd103acb38107ab1b1678cf))



# [2.23.0](https://github.com/go-courier/sqlx/compare/v2.22.0...v2.23.0)

### Features

* **feat** custom index def ([e5a928c](https://github.com/go-courier/sqlx/commit/e5a928c666d0234ecc5f1868afc7885e6477c94e))



# [2.22.0](https://github.com/go-courier/sqlx/compare/v2.21.6...v2.22.0)

### Features

* **feat** bump to golang 1.17 and deps updates ([17fcb58](https://github.com/go-courier/sqlx/commit/17fcb5811c03ca3ab2713030581586c5669115f6))



# [2.21.4](https://github.com/go-courier/sqlx/compare/v2.21.3...v2.21.4)

### Bug Fixes

* **fix** should ignore deprecated field value ([9d393eb](https://github.com/go-courier/sqlx/commit/9d393eb782d773ae2c7bce4abdfde0401770a4c1))



# [2.21.3](https://github.com/go-courier/sqlx/compare/v2.21.2...v2.21.3)

### Bug Fixes

* **fix** should ignore deprecated field value ([acdc920](https://github.com/go-courier/sqlx/commit/acdc9205aa14e11b0ef06f24226232d8d9fc2a52))



# [2.21.2](https://github.com/go-courier/sqlx/compare/v2.21.1...v2.21.2)

### Bug Fixes

* **fix** should UnwrapAll before db err check. ([6503ee0](https://github.com/go-courier/sqlx/commit/6503ee04c06eb296fcbc9c2e8b0abe5ad8ea263a))



# [2.21.1](https://github.com/go-courier/sqlx/compare/v2.21.0...v2.21.1)

### Bug Fixes

* **fix** driver connect issue when ctx pass may be cancel ([da0dbe4](https://github.com/go-courier/sqlx/commit/da0dbe4cbbdea220082fa6a6fa64a7b04edf7c22))



# [2.21.0](https://github.com/go-courier/sqlx/compare/v2.20.5...v2.21.0)

### Features

* **feat** drop logrus ([b4a9e7b](https://github.com/go-courier/sqlx/commit/b4a9e7ba17de52967d6d29064fa97cfee05a3383))



# [2.20.5](https://github.com/go-courier/sqlx/compare/v2.20.4...v2.20.5)

### Bug Fixes

* **fix(pg):** avoid gen invalid cmd ([42e911f](https://github.com/go-courier/sqlx/commit/42e911fcdfcba0f5501fa5878e0b2e9654f23fcb))



# [2.20.4](https://github.com/go-courier/sqlx/compare/v2.20.3...v2.20.4)

### Bug Fixes

* **fix(migration):** create table should dry run ([93c0b30](https://github.com/go-courier/sqlx/commit/93c0b304938fa4ef8971412891d77964f6c80f5a))



# [2.20.3](https://github.com/go-courier/sqlx/compare/v2.20.2...v2.20.3)

### Bug Fixes

* **fix(pg):** default number should be with quote and dataType ([4c0575b](https://github.com/go-courier/sqlx/commit/4c0575bf3cc9a1b03ee4715dd0fc9ffabebeb4ed))



# [2.20.2](https://github.com/go-courier/sqlx/compare/v2.20.1...v2.20.2)

### Bug Fixes

* **fix** pg comma fix ([9b52cbe](https://github.com/go-courier/sqlx/commit/9b52cbebed06597a72fe5feff207aacc1b803583))



# [2.20.1](https://github.com/go-courier/sqlx/compare/v2.20.0...v2.20.1)

### Bug Fixes

* **fix(migration):** log prev default value ([1cce629](https://github.com/go-courier/sqlx/commit/1cce629042e1877fcdb9c3f8304b00af547444cc))



# [2.20.0](https://github.com/go-courier/sqlx/compare/v2.19.1...v2.20.0)

### Features

* **feat** migration enhancement ([1fa4a9e](https://github.com/go-courier/sqlx/commit/1fa4a9e92cf79f359f3820763eeaad36b8666ea1))



# [2.19.1](https://github.com/go-courier/sqlx/compare/v2.19.0...v2.19.1)

### Bug Fixes

* **fix** enhance GetColumnName ([d7074af](https://github.com/go-courier/sqlx/commit/d7074af444c88c61f547a37636d294b18b94c0ee))



# [2.19.0](https://github.com/go-courier/sqlx/compare/v2.18.2...v2.19.0)

### Features

* **feat** alias tag for table join rename ([fc2ed04](https://github.com/go-courier/sqlx/commit/fc2ed0417fadb395d937b684bf31b1425051b924))



# [2.18.2](https://github.com/go-courier/sqlx/compare/v2.18.1...v2.18.2)

### Bug Fixes

* **fix(builder):** fieldValue fix ([ab91aec](https://github.com/go-courier/sqlx/commit/ab91aecbb684be35ec597eae670759039ac70c6b))



# [2.18.1](https://github.com/go-courier/sqlx/compare/v2.18.0...v2.18.1)

### Bug Fixes

* **fix(builder):** only struct could be visited ([ef21879](https://github.com/go-courier/sqlx/commit/ef218797c8245462060a6fd45b4bf6cc84e81cd2))



# [2.18.0](https://github.com/go-courier/sqlx/compare/v2.17.4...v2.18.0)

### Features

* **feat** support scan joined values ([f160fe4](https://github.com/go-courier/sqlx/commit/f160fe4750d488e50566c475c88069e2c1de652b))
