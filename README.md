# hexlet-path-size

[![CI](https://github.com/krenar-rm/go-test-project-242/actions/workflows/ci.yml/badge.svg)](https://github.com/krenar-rm/go-test-project-242/actions/workflows/ci.yml)
![Coverage](.github/badges/coverage.svg)

[![asciicast](https://asciinema.org/a/Vo0Iixk524cQ80yw.svg)](https://asciinema.org/a/Vo0Iixk524cQ80yw)

CLI-утилита для подсчёта размера файла или директории.

Поддерживает рекурсивный обход, скрытые файлы и человекочитаемый формат вывода.

## Сборка

```bash
make build
```

## Использование

```bash
./bin/hexlet-path-size README.md
./bin/hexlet-path-size -r -H .
./bin/hexlet-path-size -r -H -a /some/path
```

## Разработка

```bash
make test
make lint
```
