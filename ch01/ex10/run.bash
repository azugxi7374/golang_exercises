# wikipedia Go言語

echo '-> wikigo1.txt'
go run main.go --dump "https://ja.wikipedia.org/wiki/Go_(%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9E)" wikigo1.txt

echo ''
echo '-> wikigo2.txt'
go run main.go --dump "https://ja.wikipedia.org/wiki/Go_(%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9E)" wikigo2.txt

echo ''
echo 'diff wikigo1.txt wikigo2.txt: '
diff wikigo1.txt wikigo2.txt
