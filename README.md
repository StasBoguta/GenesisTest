# GenesisTest by Stanislav Bohuta

## Опис ендпоінтів
Як зазначено в завданні, дане API має три ендпоінти: /user/create, /user/login, /btcRate

### /user/create
Використовується для реєстрацї нових юзерів. Метод запиту **POST**, у body запиту необхідно передати набір параметрів: login, password. Логіном може бути лише пошта, 
а в якомті паролю використовується не порожня стрічка.

### /user/login
Використовується для авторизації юзерів. Метод запиту **POST**, у body у body запиту необхідно передати набір параметрів: login, password. В разі вдалої авторизації у 
cookies буде доданий токен авторизації.

### /btcRate
Використовується для перегляду поточного курсу біткоіну у гривні для авторизованих користувачів. Метод запиту **GET**, також cookies запcreиту повинен містити токен,
який видається користувачу після вдалої авторизації.


## Логіка роботи програми

### Збереження даних користувачів
Дані користувачів зберігаються у файлі **credentials.json** у якості набору json об'єктів. Логіни та паролі перед збереженням шифруються у base64 та зберігаються у
зашифрованому вигляді.

### Реєстрація та авторизація
Під час реєстрації (/user/create) програма перевіряє логін, який прийшов у запиті із тими, що є у файлі credentials.json. Якщо файл містить логін, то повертається відповідь
із статус-кодом 409 та повідомлення, що такий юзер вже існує. Якщо юзер з даним логіном ще не зареєстрований, то повертається відповідь із статус-кодом 200 та повідомленням,
що юзер зареєстрований.

Під час логіну (/user/login) програма перевіряє чи містить файл credentials.json логін, який прийшов у запиті. Якщо файл містить такий логін та пароль, то генерується
токен доступу та надсилається у cookies користувача. Якщо логін або пароль не співпали, то повертається 401 статус-код.

### Перегляд курсу біткоіна
Якщо запит містить токен авторизації, то програма його перевіряє. Якщо токен ще валідний (він обмежений по часу дії), то повертається відповідь із повідомленням про поточний
курс біткоїна. Якщо токена немає, або він не валідний, то повертається відповідь із 401 статус-кодом

## Курс біткоіну
Курс біткоіну береться із відкритого API біржі Kuna https://docs.kuna.io/docs/getting-started .
У програмі використовується запит GET на url https://api.kuna.io/v3/exchange-rates/bt