# Абракадабра
    4сек, 256Мб

Начало условия: Недавно Кузя разбирал чердак на даче своей бабушки и нашел очень древнюю и непонятную книгу. Кузя сразу понял, что это книга с заклинаниями, а бабушка — волшебница, просто это скрывает.

А раз бабушка умеет творить магию, то и Кузя сможет! Кузя решил тут же применить свои недюжинные навыки чтения и прочитать какие-нибудь заклинания из книги.

В дальнейшем будем считать, что все записи в книге представляют собой одну большую строку S. Все символы в книге представляют собой малые латинские буквы.

Кузя смотрел много фильмов про волшебников, поэтому знает два важных правила:

Если в данный момент прочитан символ на позиции i, то следующим Кузя должен прочитать символ на позиции pi;
Пусть ki — порядковый номер в алфавите символа на i-й позиции в тексте (a — 0-й, z — 25-й). Если Кузя за время одного заклинания должен прочитать символ на i-й позиции в mi-й раз, то вместо этого он вслух произносит символ с порядковым номером в алфавите (ki+(mi−1)⋅di)mod26.

Подробный пример находится в примечании к тестовым примерам (в самом низу).

Обратите внимание, что изменения символов при прочтении действуют только в рамках одного заклинания (mi считаются независимо для каждого прочтения заклинания).

Кузя считает, что сила прочитанного заклинания равна количеству уникальных символов, которые в него вошли. К примеру, в заклинании «zbacbef» ровно 6 уникальных символов [a,b,c,e,f,z].

Кузя нашел на обложке книги число K и понял, что для оптимального эффекта необходимо прочесть заклинания всех длин от 1 до K включительно, начав по очереди с каждого символа от 1 до N (всего Кузя планирует прочесть N×K заклинаний).

Кузя боится слишком мощных выбросов магической энергии. Поэтому он просит вас, как победителя викторины по Гарри Поттеру в 5-м классе, заранее вычислить суммарную силу всех N×K заклинаний, которые он собирается прочесть.

## Формат ввода
В первой строке дано два целых числа N и K(1≤N≤105,1≤K≤109) — количество символов в тексте книги.

Во второй строке дана строка S длины N, состоящая из малых латинских букв (Si∈[a…z]) — текст книги с заклинаниями.

В третьей строке дано N целых чисел pi через пробел (1≤pi≤N) — позиция символа, который следует прочесть после чтения символа на i-й позиции.

В четвертой строке дано N целых чисел di через пробел (0≤di≤25) — сдвиг при повторном чтении символа на i-й позиции (читайте условие и примечание).

## Формат вывода
В единственной строке выведите единственное число — суммарную силу всех N×K заклинаний, начинающихся в каждой из позиций 1,2,…,N и имеющих длину 1,2,…,K.

    Пример 1
    Ввод:
    3 7
    abz
    3 1 2
    4 0 3
    Вывод:
    74

    Пример 2
    Ввод:
    4 6
    abcd
    2 3 1 4
    1 0 2 13
    Вывод:
    62

    Пример 3
    Ввод:
    10 1000000000
    qwertzxcvb
    2 3 4 5 3 4 8 7 10 2
    1 2 3 4 5 6 7 8 9 10
    Вывод:
    259999995297