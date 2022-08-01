package tests

//func TestSearchBeneficiaries(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	mockDb := mocks.NewMockUserRepository(ctrl)
//
//	r := &api.HTTPHandler{
//		UserService: mockDb,
//	}
//
//	router := server.SetupRouter(r, mockDb)
//
//	user := models.User{
//		Model:    models.Model{},
//		FullName: "Orji Cecilia",
//		Email:    "cecilia.orji@decagon.dev",
//		Location: "ETP",
//		Password: "cece",
//		IsActive: false,
//		Status:   "active",
//		Avatar:   "image.png",
//	}
//
//	kitchenStaff := models.KitchenStaff{
//		User: user,
//	}
//	beneficiary := models.FoodBeneficiary{
//		User:  user,
//		Stack: "Golang",
//	}
//	foodBeneficiary := []models.FoodBeneficiary{
//		beneficiary,
//	}
//
//	bytes, _ := json.Marshal(user)
//
//	secret := os.Getenv("JWT_SECRET")
//	accessClaims, _ := middleware.GenerateClaims(kitchenStaff.Email)
//	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
//
//	t.Run("testing bad request", func(t *testing.T) {
//		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
//		mockDb.EXPECT().FindKitchenStaffByEmail(kitchenStaff.Email).Return(&kitchenStaff, nil)
//		mockDb.EXPECT().SearchFoodBeneficiary(beneficiary.Stack).Return(nil, errors.New("record not found"))
//		rw := httptest.NewRecorder()
//		req, _ := http.NewRequest(http.MethodGet, "/api/v1/staff/searchbeneficiary/python", strings.NewReader(string(bytes)))
//		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
//		router.ServeHTTP(rw, req)
//		assert.Equal(t, http.StatusInternalServerError, rw.Code)
//		assert.Contains(t, rw.Body.String(), "internal server error")
//	})
//
//	t.Run("testing successful search", func(t *testing.T) {
//		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
//		mockDb.EXPECT().FindKitchenStaffByEmail(kitchenStaff.Email).Return(&kitchenStaff, nil)
//		mockDb.EXPECT().SearchFoodBeneficiary(beneficiary.FullName).Return(foodBeneficiary, nil)
//		rw := httptest.NewRecorder()
//		req, _ := http.NewRequest(http.MethodGet, "/api/v1/staff/searchbeneficiary?full_name=Orji Cecilia", strings.NewReader(string(bytes)))
//		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
//		router.ServeHTTP(rw, req)
//		assert.Equal(t, http.StatusOK, rw.Code)
//		assert.Contains(t, rw.Body.String(), "information gotten")
//	})
//}
