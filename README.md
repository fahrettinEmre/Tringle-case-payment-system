İlk önce docker ile build alınması gerekmektedir.
örn:

- docker build -t fahrettin:case .

Daha sonra run edilirse uç noktalar ayağa kalkar.

- docker run -p 5050:5050 fahrettin:case

Örnek postman istekleri olup postman'e import ederek kullanabilirsiniz.

Account Create, Account Info, Payment, Deposit, Withdraw ve Transaction History uç noktaları 
, implemente edilmiştir.
