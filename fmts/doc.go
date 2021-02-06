// Code generated by 'go generate'
// source: fmts/gendoc.go
// DO NOT EDIT!
// Please run '$ go generate ./...' instead to update this file

// Package fmts holds defined errorformats.
//
// Defined formats:
// 
// 	ansible
// 		ansible-lint	(ansible-lint -p playbook.yml) Checks playbooks for practices and behaviour that could potentially be improved - https://github.com/ansible/ansible-lint
// 	csharp
// 		dotnet	(dotnet build -clp:NoSummary -p:GenerateFullPaths=true --no-incremental --nologo -v q) .NET Core CLI - https://docs.microsoft.com/en-us/dotnet/core/tools/
// 		msbuild	(msbuild /property:GenerateFullPaths=true /nologo /v:q) Microsoft Build Engine - https://docs.microsoft.com/en-us/visualstudio/msbuild/msbuild
// 	css
// 		stylelint	A mighty modern CSS linter - https://github.com/stylelint/stylelint
// 	env
// 		dotenv-linter	Lightning-fast linter for .env files. Written in Rust - https://github.com/dotenv-linter/dotenv-linter
// 	go
// 		go-consistent	Source code analyzer that helps you to make your Go programs more consistent - https://github.com/quasilyte/go-consistent
// 		golangci-lint	(golangci-lint run --out-format=line-number) GolangCI-Lint is a linters aggregator. - https://github.com/golangci/golangci-lint
// 		golint	linter for Go source code - https://github.com/golang/lint
// 		gosec	(gosec -fmt=golint) Golang Security Checker - https://github.com/securego/gosec
// 		govet	Vet examines Go source code and reports suspicious problems - https://golang.org/cmd/vet/
// 		staticcheck	Golang Static Analysis - https://staticcheck.io
// 	haml
// 		haml-lint	Tool for writing clean and consistent HAML - https://github.com/sds/haml-lint
// 	haskell
// 		hlint	Linter for Haskell source code - https://github.com/ndmitchell/hlint
// 	javascript
// 		eslint	(eslint [-f stylish]) A fully pluggable tool for identifying and reporting on patterns in JavaScript - https://github.com/eslint/eslint
// 		eslint-compact	(eslint -f compact) A fully pluggable tool for identifying and reporting on patterns in JavaScript - https://github.com/eslint/eslint
// 		standardjs	(standard) JavaScript style guide, linter, and formatter - https://github.com/standard/standard
// 	markdown
// 		remark-lint	Tool for writing clean and consistent markdown code - https://github.com/remarkjs/remark-lint
// 	php
// 		phpstan	(phpstan --error-format=raw) PHP Static Analysis Tool - discover bugs in your code without running it! - https://github.com/phpstan/phpstan
// 		psalm	(psalm --output-format=text) Psalm is a static analysis tool for finding errors in PHP - https://github.com/vimeo/psalm
// 	puppet
// 		puppet-lint	Check that your Puppet manifests conform to the style guide - https://github.com/rodjek/puppet-lint
// 	python
// 		black	A uncompromising Python code formatter - https://github.com/psf/black
// 		flake8	Tool for python style guide enforcement - https://flake8.pycqa.org/
// 		pep8	Python style guide checker - https://pypi.python.org/pypi/pep8
// 	ruby
// 		brakeman	(brakeman --quiet --format tabs) A static analysis security vulnerability scanner for Ruby on Rails applications - https://github.com/presidentbeef/brakeman
// 		fasterer	Speed improvements suggester - https://github.com/DamirSvrtan/fasterer
// 		reek	(reek --single-line) Code smell detector for Ruby - https://github.com/troessner/reek
// 		rubocop	A Ruby static code analyzer, based on the community Ruby style guide - https://github.com/rubocop-hq/rubocop
// 		sorbet	A fast, powerful type checker designed for Ruby - https://github.com/sorbet/sorbet
// 		standardrb	(standard) Ruby style guide, linter, and formatter - https://github.com/testdouble/standard
// 	rust
// 		cargo-check	(cargo check -q --message-format=short) Check a local package and all of its dependencies for errors - https://github.com/rust-lang/cargo
// 		clippy	(cargo clippy -q --message-format=short) A bunch of lints to catch common mistakes and improve your Rust code - https://github.com/rust-lang/rust-clippy
// 	scala
// 		sbt	the interactive build tool - http://www.scala-sbt.org/
// 		sbt-scalastyle	Scalastyle - SBT plugin - http://www.scalastyle.org/sbt.html
// 		scalac	Scala compiler - http://www.scala-lang.org/
// 		scalastyle	Scalastyle - Command line - http://www.scalastyle.org/command-line.html
// 	slim
// 		slim-lint	Tool to help keep your Slim files clean and readable - https://github.com/sds/slim-lint
// 	typescript
// 		tsc	TypeScript compiler - https://www.typescriptlang.org/
// 		tslint	An extensible linter for the TypeScript language - https://github.com/palantir/tslint
// 	yaml
// 		yamllint	(yamllint -f parsable) A linter for YAML files - https://github.com/adrienverge/yamllint
package fmts
