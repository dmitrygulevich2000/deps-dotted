## deps-dotted

Преобразует вывод команды `go mod graph` в `.dot` формат для визуализации графа. 
При этом фильтрует граф, оставляя в нём пути только до заданной зависимости.

Аргументы командной строки:
```
-mod название корневого модуля
-dep название модуля-зависимости
-ver semver-условие для версии зависимости, например "<2.6.1"
```

Установка:
```
go install github.com/dmitrygulevich2000/deps-dotted@latest
```

Пример использования:
```
# запуск изнутри интересующего модуля
go mod graph | deps-dotted -mod=github.com/you/your-module -dep=github.com/stretchr/testify -ver="<=1.5" | dot -Tsvg -o deps.svg
```
