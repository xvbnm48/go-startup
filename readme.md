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


tes uji coba login
	//user, err := userService.Login(input)
	//if err != nil {
	//	fmt.Println("terjadi kesalahan")
	//	fmt.Println(err.Error())
	//}

	//fmt.Println(user.Email)
	//fmt.Println(user.Name)

