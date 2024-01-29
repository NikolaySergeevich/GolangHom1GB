## Пояснения

Условно путь будет указываться через "/". Я могу разбить весь путь через **split** и взять крайний аргумент, что бы его уже разбивать на название и расширение. Обычно путь какой-либо не занимает много строк текста, как например рассказ чей-то. Поэтому разбив его на части, мы получаем полное название файла - **lastSplit**

Так как перед любым расширением ставится точка, то дальше проверяю есть ли точка в названии файла. Если нет, то говорю, что **lastSplit** и есть имя файла. А сам файл без расширения. Иначе, код впадает в цикл и проходится по названию файла.

Используя метод **strings.LastIndex()** мы находим самую крайнюю точку в названии файла и всё то, что лежит по левую сторону от неё записывается в имя файла, а всё то, что по правую в расширение.

--------------------------------------------------------------------------------------------------------------------------------------------------

# Дополнение, которое я не заметил сразу.

Когда разбили сплитом и нашли полное имя файла с расширением, то не нужно for и по одному элементу добавлять в переменную типа стрингбилдер.

Просто воспользоваться стрезом и в переменную fileName записать от начала и до индекса с точкой, а в перменную fileExt от точки и до конца
