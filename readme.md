# Catatan go startup


folder repository itu gunanya untuk sebagai tempat untuk terkoneksi ke database

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

