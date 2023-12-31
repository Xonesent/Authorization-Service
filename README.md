Продолжение работы с сервисами на go. В данном проекте я реализовал авторизацию и поставил заглушку api из своего предыдущего репозитория.  

# Об авторизации  
Перед тем как перейти к функционалу сервиса по сокращению ссылок, поддерживающий два http запроса:  
1. Post - сохраняет оригинальный URL в базе и возвращает сокращённый  
2. Get - принимает сокращённый URL и возвращает оригинальный URL, если он есть в базе  
Пользователю необходимо зарегестрироваться:  
1. Post - Sign Up - получение логина и хэш пароля в базу  
2. Post - Sign In - получение jwt токена для аутентификации в сервисе и получения доступа к функционалу на 12 часов  

Проект был сделан для получения практического опыта и в нем реализована:  
- Следование REST API дизайну  
- Работа с фремворком gin  
- Работа с базой данных postgresql  
- Работа с окружением godotenv + viper  

# Работа с авторизацией:  
Post - Sign Up  
Запрос: curl.exe -X POST localhost:8080/auth/sign-up -H "Content-Type: application/json" -d '{"login" : "Xonesent", "password" : "test"}'  
Ответ: {"id": 1}  

![Alt text](sign-up.png)  

Post - Sign In  
Запрос: curl.exe -X POST localhost:8080/auth/sign-in -H "Content-Type: application/json" -d '{"login" : "Xonesent", "password" : "test"}'  
Ответ: {"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQwNzg5NjAsImlhdCI6MTcwNDAzNTc2MCwidXNlcl9pZCI6MX0.ypxpk6sNjOpclzsU4BmPBKx9L1YjnfJ_z2eQ1zkrW20"}  

![Alt text](sigh-in.png)  

Post - Send Url - Дальнейший доступ к функционалу предоставляется с помощью Bearer Token  
Запрос: curl.exe -X POST -H "Content-Type: application/json" -H 'Authorization: Bearer <...>' -d '{"login" : "Xonesent", "password" : "test"}' localhost:8080/auth/sign-in  

![Alt text](bearer.png)  

# Что хотелось бы еще сделать  
- Авторизацию через письмо на почте  
- Реализацию большего количество http методов  
- Запуск из docker-compose  
- Unit тестирование  

Буду рад получить фидбэк / советы!  
Почта - 1pyankov.d.s@gmail.com  
Телега - @Xonesent  