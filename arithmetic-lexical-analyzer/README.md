Необходимо создать консольную утилиту, которая на вход принимает текстовый файл с
последовательностью символов, представляющих произвольное арифметическое
выражение, а на выходе создаёт файл с последовательностью токенов и информацией по
ним и файл, содержащий таблицу символов.

Пример работы утилиты: lab2.exe inputExpr.txt tokens.txt symbols.txt

Утилита принимает в качестве параметра имя файла с выражением inputExpr.txt:

var1 + (9.5 – 5 * (var2 - 0.6)) / var3

В выражении могут быть четыре арифметические операции ( +, -, *, / ) скобки любого
уровня вложенности, вещественные и целочисленные константы, идентификаторы
переменных с традиционными правилами именования (состоят из букв английского
алфавита, цифр, знака нижнее подчёркивание _ , должен начинаться с буквы или знака
нижнего подчёркивания).

Далее, произведя лексический анализ, утилита создаёт выходной файл с
последовательностью токенов tokens.txt следующего содержания:

* <id,1> - идентификатор с именем var1
* <+> - операция сложения
* <(> - открывающая скобка
* <9.5> - константа вещественного типа
* <5> - константа целого типа
* <*> - операция умножения
* <(> - открывающая скобка
* <id,2> - идентификатор с именем var2
* <-> - операция вычитания
* <0.6> - константа вещественного типа
* <)> - закрывающая скобка
* <)> - закрывающая скобка
* </> - операция деления
* <id,3> - идентификатор с именем var3

также создаётся файл с таблицей символов symbols.txt:
* 1 – var1
* 2 – var2
* 3 – var3

В процессе работы лексического анализатора могут возникать ошибки анализа из-за
недопустимых символов в потоке. Подобные ошибки должны выявляться и правильно
обрабатываться, выдавая в консоль сообщение об возникшей ошибке и позиции в потоке
символов, на котором она возникла, например:

_Лексическая ошибка! Недопустимый символ “#” на позиции 4_

_Лексическая ошибка! Идентификатор «1var» не может начинаться с цифры на позиции 15_

_Лексическая ошибка! Неправильно задана константа «0.5.4555» на позиции 9_

Обратите внимание, синтаксические ошибки лексическим анализатором не
обрабатываются! Эта задача синтаксического анализатора. Таким образом ошибки вида
«не закрыта открывающая скобка» или «два знака операции следуют подряд» или «у
оператора не хватает операнда» не определяются на этапе лексического анализа.