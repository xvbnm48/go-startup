# Catatan go startupppppp 2022
!!!!!


folder repository itu gunanya untuk sebagai tempat untuk terkoneksi ke database
aaaaaaaaaaaaaaa
## urutan folder 

1. request
2. handler
3. service 
4. repository
5. DB


rabu 24-11-2021 belajar docker
kamis belajar docker


tes uji email di temukan
```go
	// uji coba
	userByEmail, err := userRepository.FindByEmail("sakichan@challange.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	if userByEmail.ID == 0 {
		fmt.Println("user tidak ditemukan")
	} else {
		fmt.Println(userByEmail.Name)
	}
```


tes uji coba login
```go
	user, err := userService.Login(input)
	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)
```

di gorm untuk mendapatkan sebuah file dari header itu dengan menggunakan 
```go
c.FormFile
```
dengan parameternya itu nama headernya

terus untuk untuk menyimpan gambarnya dengan perintah
```go
c.SaveUploadedFile()
```
dengan paramerternya , file tadi yang ditangkap dan path nya

untuk tes token
```go
	fmt.Println(authService.GenerateToken(10))

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxN30.9AhPIz7AyPWBJyk0sb6uFuN2QrPctGRr45n1l7jxaL4")
	if err != nil {
		fmt.Println("error")
	}

	if token.Valid {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}
```
untuk tes ambil data foto
```go
	campaings, err := campaignRepository.FindByUserID(1)

	fmt.Println(len(campaings))

	for _, campaign := range campaings {
		fmt.Println(campaign.Name)
		if len(campaign.CampaignImages) > 0 {
			fmt.Println(campaign.CampaignImages[0].FileName)

		}
	}
```

untuk simpan foto tes
	userService.SaveAvatar(1, "images/avatar.jpg")

# untuk mengetes service create campaign 
``` go
input := campaign.CreateCampaignInput{}
	input.Name = "penggalanga  dana JKT48 agar tidak bubar"
	input.ShortDescription = "macet"
	input.Description = "macet cokkkkkk"
	input.GoalAmount = 1000000000
	input.Perks = "hadiah,hadiah2,hadiah3"

	inputUser, _ := userService.GetUserByID(1)
	input.User = inputUser

	_, err = campaignService.CreateCampaign(input)
	if err != nil {
		log.Fatal(err.Error())
	}
```


## cara test service

```go
user, _ := userService.GetUserByID(18)
	input := transaction.CreateTransactionInput{
		CampaignID: 18,
		Amount:     92000000,
		User:       user,
	}

	transactionService.CreateTransaction(input)

```



sessions localhost:8080/api/v1/sessions
input
```json
{
	"email" : "akisuda@tokisen.com",
	"password": "12345"
}
```
response 
```json
{
	"Meta": {
		"Message": "successfuly login",
		"Code": 200,
		"Status": "success"
	},
	"Data": {
		"id": 17,
		"name": "aki suda",
		"occupation": "idol",
		"email": "akisuda@tokisen.com",
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxN30.SVFwfK26LsCByEIbngH4r8TZ94mADgajGhP8ljZjn9g"
	}
}
```
new user localhost:8080/api/v1/users
input
```json
{
	"name" : "misaki chan",
	"email" : "misaki@kawaii.com",
	"password" : "12345",
	"occupation" :"idol"
}
```

response 
```json
{
	"Meta": {
		"Message": "account has been registered",
		"Code": 200,
		"Status": "success"
	},
	"Data": {
		"id": 21,
		"name": "misaki chan",
		"occupation": "idol",
		"email": "misaki@kawaii.com",
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyMX0.maqABc0-BTgrK6rXKQmX32hbBZyqIZLy2IE9ofQtkvM"
	}
}
```

email check localhost:8080/api/v1/email_checkers

input
```json
{
	"email" : "akarin@kawaii.com"
}
```
response
```json
{
	"Meta": {
		"Message": "email is registered",
		"Code": 200,
		"Status": "success"
	},
	"Data": {
		"is_available": false
	}
}
```

upload avatar localhost:8080/api/v1/avatars
input 
multipart , value avatar
authorization ,Bearer token

response
```json
{
	"Meta": {
		"Message": "avatar image success uploaded",
		"Code": 200,
		"Status": "success"
	},
	"Data": {
		"is_uploaded": true
	}
}
```

get campaings
input
kosong
response 
```json
{
	"Meta": {
		"Message": "LIst Of Campaigns campaigns",
		"Code": 200,
		"Status": "success"
	},
	"Data": [
		{
			"id": 1,
			"user_id": 0,
			"name": "campaign",
			"short_description": "shorttt ",
			"image_url": "campaign-images/fotoq.jpg",
			"goal_amount": 1000000,
			"current_amount": 0,
			"slug": "campaign1"
		},
		{
			"id": 2,
			"user_id": 0,
			"name": "campaign2",
			"short_description": "shorttt ",
			"image_url": "",
			"goal_amount": 1000000,
			"current_amount": 0,
			"slug": "campaign2"
		},
		{
			"id": 3,
			"user_id": 0,
			"name": "campaign2",
			"short_description": "shorttt ",
			"image_url": "",
			"goal_amount": 1000000,
			"current_amount": 0,
			"slug": "campaign3"
		}
	]
}
```

get images localhost:8080/images/1-1068795.png
input kosong 
response
```json
gambar.jpg
```

get campaign satu aja localhost:8080/api/v1/campaigns/3

response 
```json
{
	"Meta": {
		"Message": "campaign detail",
		"Code": 200,
		"Status": "success"
	},
	"Data": {
		"id": 3,
		"name": "campaign2",
		"short_description": "shorttt ",
		"description": "longg",
		"image_url": "",
		"goal_amount": 2000000000,
		"current_amount": 0,
		"user_id": 19,
		"slug": "campaign3",
		"perks": [
			"prtks1",
			"perks2",
			"perks3"
		]
	}
}
```

update campaign localhost:8080/api/v1/campaigns/6

input 
```json
{
	"name":"konser cho tokimeki sendenbu indonesia 48",
	"short_description": "konser tokimeki sendenbu",
	"description":"konser tokimeki sendenbu di indonesia", 
	"goal_amount":2000000,
	"perks" :"lightstick, t-shirt, towel, photopack"
}
```

Authorization, Bearer token

response 
```json
{
	"Meta": {
		"Message": "Failed to update Campaign",
		"Code": 400,
		"Status": "error"
	},
	"Data": null
}
```

upload campaign images localhost:8080/api/v1/campaign-images
input
multipart 
campaig_id
is_primary
file

response
```json
{
	"Meta": {
		"Message": "campaign image success uploaded",
		"Code": 200,
		"Status": "success"
	},
	"Data": {
		"is_uploaded": true
	}
}
```

get transaction campaign
localhost:8080/api/v1/campaigns/6/transactions

response
```json
{
	"Meta": {
		"Message": "Success to get campaign Transaction",
		"Code": 200,
		"Status": "success"
	},
	"Data": [
		{
			"ID": 1,
			"CampaignID": 6,
			"UserID": 17,
			"Amount": 1000000,
			"Status": "paid",
			"Code": "",
			"User": {
				"ID": 17,
				"Name": "aki suda",
				"Occupation": "idol",
				"Email": "akisuda@tokisen.com",
				"PasswordHash": "$2a$04$VupfIutRg2RkxC4KwJN/e.JUnSsdKycrGQTNx54xYQhskP/KkVu3K",
				"AvatarFileName": "",
				"Role": "user",
				"CreatedAt": "2021-11-28T11:56:06+07:00",
				"UpdatedAt": "2021-11-28T11:56:06+07:00"
			},
			"CreatedAt": "2021-12-15T17:21:20+07:00",
			"UpdatedAt": "2021-12-15T17:21:20+07:00"
		}
	]
}
```
get user transaction

response 
```json
{
	"Meta": {
		"Message": "Success to get user Transaction",
		"Code": 200,
		"Status": "Success"
	},
	"Data": []
}
```

create transaction
localhost:8080/api/v1/transactions

input 

```json
{
	"campaign_id" : 7,
	"amount" : 100000
}
```

response 
```json
{
	"Meta": {
		"Message": "Success to create New Transaction",
		"Code": 200,
		"Status": "Success"
	},
	"Data": {
		"id": 11,
		"campaign_id": 7,
		"user_id": 17,
		"amount": 100000,
		"status": "Pending",
		"code": "",
		"payment_url": "https://app.sandbox.midtrans.com/snap/v2/vtweb/7c43de0d-9d02-44ee-b33f-a2c1734984a3"
	}
}

```

create campaign localhost:8080/api/v1/campaigns

input 
```json
{
	"name" :"tokisen konser in indonesia",
	"short_description" :"konser tokisen",
	"description":"konser yang dilaksanakn di jakarta indonesia",
	"goal_amount": 1000000,
	"perks" : "t-shirt,lightstick,towel"
}

```

response 
```json
{
	"Meta": {
		"Message": "success to created campaign ",
		"Code": 200,
		"Status": "success"
	},
	"Data": {
		"id": 7,
		"user_id": 17,
		"name": "tokisen konser in indonesia",
		"short_description": "konser tokisen",
		"image_url": "",
		"goal_amount": 1000000,
		"current_amount": 0,
		"slug": "tokisen-konser-in-indonesia-17"
	}
}
```
