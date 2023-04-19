

# ScamfariHelper
Сonsole application to make life easier for hunters of Scamfari

# Instruction
1. Download latest release in [Releases](https://github.com/Sanichka/ScamfariHelper/releases)
2. Extract all files from ScamfariHelperVx.x.x.zip in any folder.
3. Add your reported wallets in wallets.txt file.
4. Additionally you can change new, and old wav files in cfg.env for your unique alert sounds. Or specify your txt file in config.
5. Start ScamfariHelper.exe and wait till the console window will show up.
6. Start searching wallets, once you copy anything longer than 30 character Scamfari helper will play audio based if the copied text is already in wallets.txt.

# Інструкція

1. Завантажте останні реліз у вкладці [Releases](https://github.com/Sanichka/ScamfariHelper/releases)
2. Розпакуйте всі файли зі ScamfariHelperVx.x.x.zip у будь-яку папку.
3. Додайте свої гаманці у файл wallets.txt.
4. Крім того, ви можете змінити аудіо файли wav у cfg.env на ващі унікальні звуки для сповіщень. Або вказати свій текстовий  файл у конфігурації.
5. Запустіть ScamfariHelper.exe та зачекайте, доки з’явиться вікно консолі.
6. Почніть пошук гаманців як тільки ви скопіюєте щось довше ніж 30 символів, помічник Scamfari відтворить аудіо на основі, того чи є скопійований текст у wallets.txt.

# Contact details
For any questions please use the official [Telegram Scamfari group](https://t.me/scamfari_public)  
For any bugs, preferably check your latest .log file for errors and ensure all files are present and the cfg.env file is filled with filenames. If you still encounter any bug or error, please fill out a [new issue](https://github.com/Sanichka/ScamfariHelper/issues/new/choose). List of known issues you can find here [GitHub Issues](https://github.com/Sanichka/ScamfariHelper/issues).
# Контактна інформація
З будь-якими запитаннями звертайтесь в офіційну [групу Telegram Scamfari](https://t.me/scamfari_public)  
При виявленні будь-яких помилок перш за все перевірте останній .log файл на наявність помилок і переконайтеся, що всі файли присутні і файл cfg.env заповнено іменами файлів. Якщо ви все ще стикаєтеся з будь-якою помилкою, заповніть [new issue](https://github.com/Sanichka/ScamfariHelper/issues/new/choose). Список відомих помилок [GitHub Issues](https://github.com/Sanichka/ScamfariHelper/issues).

# Інструкцію пошуку гаманців
Почнемо з того що потрбно використовуваи впн, тому обирайте і встановлюйте собі будь-який барузерний або декстопний. Тепер перейдемо до основ. Використання логічних виразів і фільтрів в пошукових системах(гугл, яндек, duckduckgo)  є найефективнішим варіантом пошуку. Кожна з пошукових систем використовує своїх ботів/crawler/web spider щоб "ходити" по сайтах і індексувати контент згідно налаштувань сайту і по ним видавати вам результати. Повний список операторів для пошуку в гуглі можете подивитись [тут](https://ahrefs.com/blog/google-advanced-search-operators/), ми будемо використовувати тільки основні.
- AND логічна умові і
- OR логічна умова або
- "" щоб вказати що слово обов'язково має бути на сторінці
- site:vk.com пошук на певному сайті
- знак - перед словом або іншим оператором чи -site:vk.com, щоб виключити слова/сайти з результатів
- after:2022-02-24 щоб вказувати після якої дати має бути результат(наприклад шукати свіжі гаманці)
- daterange:22055-23055 пошук в проміжку часу(у [Julian date format](http://www.longpelaexpertise.com/toolsJulian.php))
- Використання дати з меню гугл Інструменти - за період чи певний час <a href="https://ibb.co/bKymv13"><img src="https://i.ibb.co/mcYTyJB/image.png" alt="image" border="0"></a>

Список криптовалют/мереж які ми шукаємо і вони підходять для репорту:
- NEAR
- AURORA
- SOL / PSL
- BNB / BEP20 / BEP2 / BSC
- BTC
- Polygon MATIC
- OK / Okcash
- TRON / TRX / TRC20

Розбір прикладу пошукового запиту гугл: **after:2022-02-24 "СВО" AND (ETH OR BTC OR BCH OR eth OR btc OR bch) AND -урожай AND -блендер AND -миксер AND "реквизиты"**
В цьому запиті ми шукаємо всі сторінки які з'явилися в пошуку після початку повномасштабної війни вказавши ключове слово after:, щоб обв'язково було слово СВО і групуємо список криптовалют дужками і логічними операторами OR, тобто хоча б одна з них була на сторінці. Додатково їх можна зробити обов'язковими взявши їх в "". І прибираємо сторінки на яких можуть бути ці абревіатури через знак -, щоб виключити всілякі міксери, біржі і т д. А також щоб в результаті обов'язково було слово реквизиты.
Також можна використати наступний запит щоб шукати тільки на сайтах домену .ru: **after:2021-02-24 сбор AND (ETH OR BTC OR BCH OR eth OR btc OR bch) AND -урожай AND -блендер AND -миксер AND СВО**. Але результатів багато і як тренування можете спробувати додати більше операторів і "" щоб відсіяти зайві результати.
До кожного пошукового запиту можете додавати -адреса вже репортнуого гаманця, щоб вам не траплялись дублікати. Як приклад є один підсканційний сайт з новинами і ось частина мого запиту для нього: **site:southfront.org AND ("ETH" OR "BTC" OR "ERC20" OR "trc20" OR "TRC20" OR "erc20") -bc1qctv99yh0ewg6x5r9fy5e7lqm28t9rza4h4cy4k -bc1qgu58lfszcpqu6fd8l98m378wgzugyg9y93lcym -3Gbs4rjcVUtQd8p3CiFUCxPLZwRqurezRZ**.
Можете доповнити його знаками - і гаманцями що ви знайшли, щоб витягнути з нього максимум, кінцевий запит у вас вийде приблизно 456+- символів. Але скоріш за все всі їх гаманці вже є в публічній базі і вам його не зарахуют, тому просто як приклад і тренування.

Так само у [вконтакті достатньо запиту](https://vk.com/search?c%5Bexclude%5D=TBUxP3Mbx5cYUCB8B3VVZEqABdum1LXiu1&c%5Bper_page%5D=40&c%5Bq%5D=%D0%B1%D0%BE%D0%B9%D1%86%D0%B0%D0%BC%20eth%20btc&c%5Bsection%5D=statuses): **сво eth btc** або **бойцам eth btc**. І додавання вже репортутих гаманців через пробіл в поле **Исключить слова**.
Приклад запиту в яндексі, все так само але логічні оператори мають трохи інший вигляд, детальніше можете знайти в інтернеті: **date:2022**** **"сбор" & ("ETH" | "BTC" | "BCH" | "eth" | "btc" | "bch") -кран -обмен -краны -майнинг**
Також інші сайти, соц. мережі і т д мають свої системи пошуку з певними фільтрами або аналогом операторів, тому можете ще використовувати їх фішки. Для пошуку в телеграмі використовуйте сайти які "індексують"/парсять групи телеграм або мають списки каналів по категорії.
Використовуючи вказані оператори, фантазуючи які слова краще підібрати, на яких сайтах шукати і т д сподіваюсь ви знайдете нові для себе гаманці і полегшите для себе процес пошуку. Як ідея для тих хто програмує, все описане вище можна використати для софту з парсером який буде створювати лінки для пошуку і сам переходити, шукати і парсити результати з сайтів.

# Project Support
Current development time 12 hours. Donations will support and encourage project development!  
Поточний час розробки 12 годин. Пожертви підтримають та заохочуватимуть розвиток проекту!  
**[Donate](https://send.monobank.ua/jar/3kQv2UwkhF)**  
![Donate QR code](https://i.imgur.com/mL5LpwZ.jpeg)