package art

import (
	"context"
	"time"
)

type Article struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Article   string    `json:"article"`
	Comment   string    `json:"comment"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteStore interface {
	CreateNote(ctx context.Context, note Article) (string, error)
	ListByUserID(ctx context.Context, id string) ([]Article, error)
}

var input = []string{
	"B2mlQHkF_Bo",
	"B2mld_8pMNP",
	"BrjZ1v3lSo4",
	"B20deL_l-Xl",
	"B20dpj4hCUf",
	"B22iMHwgQn6",
	"B22ibDsBYu4",
	"B24Jr0vHJ4N",
	"B2mV-x4HxVa",
	"B2fQU1aBMlz",
	"B2mlWBQnVve",
	"B2mlaX5Bpej",
	"B20dOuXF7PJ",
	"B2mlkQYHrOu",
	"B2mlrR4A4pN",
	"B20RlUhBf1b",
	"B21w75TFH50",
	"B22ia91ANBF",
	"B2mlmBvnTC_",
	"B22iUa2AGkH",
	"B2mltr5gBDl",
	"B2mdu7IlDZ_",
	"B2kdBXpA3hS",
	"B2a8-4vnFMw",
	"B2ml0elho_7",
	"B2ml2qYoUUK",
	"B24Jtg0BgBj",
	"B2mlp-zpYVp",
	"B2ml52IFkMW",
	"B2mmwichKfF",
	"B2mUxnsghQp",
	"B2mmU11lPsL",
	"B22iXNCniMr",
	"B2mmfbGA1d9",
	"B30dqSbg5vf",
	"B1BrCuFos4T",
	"B2shI1LHEKG",
	"B2st-DJDSYC",
	"B58Gx_yAWZO",
	"B2s7PNNBRGo",
	"B2ml6NNgr65",
	"B2mllBFBLPz",
	"B2mmGtoBBDz",
	"B2KYGWwln9K",
	"B2ml_yLlLxd",
	"B2mmBIKHqJ4",
	"B2mmNYSBch6",
	"Bo73F_1hJOW",
	"B0_S0-0ogzo",
	"B260tnxhkBZ",
	"B2ml9bXg1rQ",
	"B2ml9zsn4Wo",
	"BiR6KwHFOtM",
	"B23nWz-olh0",
	"B2mmMXDlQ-I",
	"B2mnVYJB5KS",
	"B2ml6XIFZfW",
	"B1Bq9wYo22o",
	"B2xWMR3B4Rn",
	"B2md_P1j5qZ",
	"B2ml1lfl9rI",
	"B2lVLpYg8QV",
	"B2mmWUYnmJY",
	"B2mmixtnIB0",
	"B2l6rFUF546",
	"B1yVjDxl1ML",
	"B2mmQFzBPw1",
	"B2mmT0iBKud",
	"B2md__bDum3",
	"B2iCaFqHG6l",
	"B2meItDDVno",
	"B1Bl5WUFXVT",
	"B2mmodjFyj7",
	"B2mnTOHHT_b",
	"B2mmRSzFy_Y",
	"BRC21pRDNkc",
	"B2mmdGjFaWi",
	"B4EEUt7JTmJ",
	"B2mmlVQAd_H",
	"B2mm1JABlCg",
	"B2h6_r3hujZ",
	"B22ia-WlAJt",
	"B2mmSQzlEUE",
	"B2mmeYTD3av",
	"B2mb-RWiKma",
	"B22VstCAxlh",
	"B24kuwlF3kI",
	"B24D7KencBN",
	"B24k2reFhee",
	"B24k0Tyn_AD",
	"B24lEYKF-OJ",
	"B2mUaj4l3U8",
	"B2mmnJunQs7",
	"B20Q1-UAROe",
	"B2PVAUJhUhM",
	"B20ZcOlhCZQ",
	"B261hd5pAn0",
	"B261zxpHmSS",
	"B2uWXjFJ6oi",
	"B27eQg4JCMY",
	"B2jvz9hnS6p",
	"B2fIsQUhn6J",
	"B20d405gG04",
	"B22ihAyD93v",
	"B2dKSpBgd4J",
	"B23nbDOH_Aa",
	"B2mKz9YF1lD",
	"B20dX2pDRje",
	"B21eI2sAhSS",
	"B22jCJzgV7K",
	"B22jIR9hyAv",
	"B22jNiElHVg",
	"B23nf0zpsFo",
	"B27eXY2lNH8",
	"B2mmxk0BYwE",
	"B20dmm4F_Ty",
	"B2mUSkoHzrI",
	"B2mm_9bAxzG",
	"B2mnBiuBeng",
	"B2mnFeNHo1u",
	"B2mnM5_pK8h",
	"B2mnP3tFDLN",
	"B2mnOkCHwrA",
	"B23ng59nJgF",
	"B2mnAy0h2sY",
	"B-cBbE0opOR",
	"B2mlrOKAdnr",
	"B2mnGg8B_db",
	"B2a9jnpJwci",
	"BxbSpLvFssI",
	"B3KlPcAHmzW",
	"B4NYL2bBwpa",
	"B2mlnZ1hNU3",
	"B2mGSidjiCX",
	"B2mnSg0h_Pb",
	"B2HHFnGBDPY",
	"B2mnURbBNsf",
	"B2mnQ51pPzp",
	"B2mnYfZF--Q",
	"B2mmYMTgz62",
	"B3L1smeHzi0",
	"B2mnptGlhtn",
	"B2mW3BTBvsq",
	"B22ijWUARCq",
	"B2mmB5NiSyA",
	"B2ZvpRVg5ub",
	"B2mn2c0pkd2",
	"B20dpS4p1ug",
	"BzJ5U4aJu8A",
	"B3KlR_0Hj8c",
	"B2mnX1fh-6B",
	"B2mnfbvgGYx",
	"B2mnr8uBmwu",
	"B2mntiHnVtb",
	"B20eJgwFtKT",
	"B2mscvNlxyn",
	"B2moAMNBOXM",
	"B2mesCLl35B",
	"B20eLqznHTO",
	"B27e7IzFrPu",
	"B2ljtCvgSSU",
	"B2mnT1WlhLn",
	"B1_HqY1Jicg",
	"B20eMGjFM1L",
	"B22ij23IWYv",
	"B2mnbyBH2Tg",
	"B2mQ-JdAY-U",
	"B25jb2KFQOi",
	"B2mnb5PlzwP",
	"B3KlPcDIb0p",
	"B2Uy5ytAxEB",
	"B2aaaOvB-dL",
	"Bq2OuSKB91o",
	"B2mnq_nAr2V",
	"B2mn5NhnMd-",
	"B20eMpZAgx7",
	"B20fHXBJick",
	"B22ikU7g1ag",
	"B2mnhLWjBhL",
	"B2UyZ3zgM9Y",
	"B2mnrLLhecd",
	"Br5uMcEjaNP",
	"B2mlzTcl0Ff",
	"B20eR86BT6s",
	"B7VhKVGpwM5",
	"B2mn9w-jtRL",
	"B3KlUcvloF-",
	"B2moC8bF_3D",
	"B2moEMeAkB7",
	"B22imkwn3PD",
	"B2mo3qUh2B4",
	"B2mo683H5aL",
	"B2mpFCtnN5l",
	"B2mpHYFHeyO",
	"B3QAvLnjVh-",
	"B1L3r4egSAC",
	"B1UnFXSgtGj",
	"B2lXoDOF4EM",
	"B2moXoHltHw",
	"B2mo08hAfxJ",
	"B20eNXnlLUg",
	"B20ebC1gHKC",
	"B2wIeBthuEQ",
	"Bxiot4TpHxX",
	"B20eWyGHqgy",
	"B2mZv25FrH5",
	"B2moRVNhkJI",
	"B2moRHFApON",
	"B19maXdgXpJ",
	"B27fFpGlM3f",
	"B2moNu8HdfQ",
	"B2moIHaFlZD",
	"B2dB4AqnOkD",
	"B2mjhoTJFPZ",
	"B2moNAVFc9b",
	"B2k9Z-WlKLA",
	"B2moQR_HJmh",
	"B2moh4QFqI6",
	"B-cBVAeokqG",
	"B2McPFylVfq",
	"B2fmZZpFOVi",
	"B2moZ5MDFsN",
	"B2mofDnHcF5",
	"B2mo9CMhNA-",
	"B2mpBlVFR63",
	"B27fz8rhsWT",
	"B2moyEWFjSL",
	"B2mFt8XnZ2g",
	"B2hf-7dnViy",
	"B2Hg175H_SA",
	"B2moTj-gpSB",
	"B2mB2NTAlMy",
	"Bz0bpQBBiSZ",
	"B2morxVB44s",
	"B2movIgHU_E",
	"B2mo0i7iexD",
	"Bnj_SjIg8nI",
	"B2mo_E_g_zO",
	"B20eY4vghFn",
	"B3KlUsWHEFu",
	"B22ipx6lfhl",
	"B2x4J8KFGvd",
	"B2monvLAMel",
	"B2mn60KBKAd",
	"B2kGDGwHlcj",
	"B2mGHhWnhHq",
	"B2z9KfnBBlK",
	"B20eeVhhXQi",
	"B20enTcB1fh",
	"B2t1cuyBbpO",
	"B2mWIRcgV_r",
	"B2mn_ZPA-yz",
	"B2mopRZhMB_",
	"B2mogaTFz3z",
	"B2mpUkKnmy3",
	"B20ej70n0kO",
	"B20enqgH5X3",
	"B2mo6logw_V",
	"B2mc449lcqg",
	"B2kDJ5Hnt8-",
	"B2mpH-HgcGF",
	"B2mpDTAnQme",
	"ByoAfJ2ni6e",
	"B2mpPZHlxpN",
	"B2mzylJFHMH",
	"B2mpJPwA6kx",
	"B2m-UWrnZAD",
	"B2mpJnrpNsT",
	"B20eogaBMmw",
	"B20erd3Dz1d",
	"B20etSlhP86",
	"B2z_jMRofZ4",
	"B20ezjMBzf1",
	"B22ifh-Bvxu",
	"B24vwFwFpt5",
	"B3AEjsZB2GK",
	"B3AErQUh3Ey",
	"BtPP6N3ALzD",
	"B2mpMDxhizf",
	"B2mpSoTDhoB",
	"B194bvHheAN",
	"B20eu1Oggk6",
	"B20esQhAyov",
	"B25-hHui1hM",
	"B3KlO2fB3d2",
	"B2ml-00iW2k",
	"BwWrEEDAs80",
	"B2mAH36jQNU",
	"B2m-nUvHLK8",
	"B20erawg0le",
	"B2nBRY-Aw9I",
	"B20eSEjATg9",
	"B2mpYF3AmS0",
	"B3KlW5ApV0L",
	"B2mpffMHwqo",
	"B2mpirvpVj0",
	"B2mpmOXBVxS",
	"B2mpjN5p1Tg",
	"B2mpnxoB9nu",
	"B22jT11HHik",
	"BtTEEuPAVMc",
	"B2mpZD2jqeg",
	"B2mp24YgGRS",
	"B2mqDOoponx",
	"B20ezXvhBf0",
	"B2mqPM-FDNP",
	"B22iykggzNG",
	"B2w1MA6FQwC",
	"B2mpn7oHjL6",
	"B2mpfSOHug2",
	"B2mh-34FZVl",
	"B2k6vEzA-H6",
	"B22i06gBlof",
	"B2md69kgEb6",
	"B2nqtJKgJPe",
	"B25GPlyHbgu",
	"B22jfUJBkQF",
	"BtDpCGJgz7c",
	"B2kgs74lbQY",
	"B2xIq9bHk7t",
	"B2mps8EHNiH",
	"B2mqJz8nx-J",
	"B22i2eUFH2Q",
	"B23nlvbp7LM",
	"B2mpiI7Bljc",
	"B2mqGT4FPuJ",
	"B20Nfvulu_4",
	"B22i2giJWcO",
	"B2UyoclHvmP",
	"B3L2jUpI0FR",
	"B2mM6ENpu3K",
	"B1Z8R8NnS4R",
	"B2mp8B2nfPs",
	"B2mqEaNheyT",
	"B2mqNERJp_5",
	"B2mqByUlZcb",
	"B3KlY3gADT4",
	"B3KlaQSlP6C",
	"B3KlgK6HlV5",
	"B2mqARqgKNh",
	"B2mqIR2BxMx",
	"B2mqAVBlRl3",
	"B2mqaSMnDrX",
	"B20fGKUFilh",
	"B20fk_UBN53",
	"B22i4ralIzo",
	"B2dItPcAHwI",
	"B0dBswqiIgB",
	"B20fG-wB9u7",
	"B20gDFUBCuH",
	"B22i-E_hEa7",
	"B22jVUrgiK-",
	"B22jiFZgX2y",
	"B23nuVxCXUN",
	"B23n_11Bkj3",
	"B3KlhlOAN7r",
	"B3Kj3LVhYpz",
	"B0asZgcBvGW",
	"B2mRM0LgPf_",
	"BzvpoF-JYYg",
	"B2mkuGbhFsQ",
	"B2mrQ96FWjT",
	"B2ms6BMBZzM",
	"B25zg5rAbpw",
	"B27mzx4ndXB",
	"B3AEzxigtS0",
	"B2mqZaVBJ36",
	"B2mqaXMJVVa",
	"B2lM_mlFFkW",
	"B2mZQbRBrfU",
	"B2mqikOJL_Y",
	"B2mqk3eJWd1",
	"BzulPr_ni8K",
	"B2VmjG7A3HK",
	"B2mqjuaFJv8",
	"B20e77ygX-M",
	"B2zD90ugftM",
	"B2mqo1zF3Ni",
	"B20f4a_h7G1",
	"B2zPWKQnFeg",
	"B25TzttHYhg",
	"B25-rsIBqnj",
	"B2mrTN5Fksh",
	"B2mqnuxJXt-",
	"B2mquuBhAgc",
	"B2mqzWlDnNw",
	"B20e8TmAgZp",
	"B2mrLgzJfZd",
	"B2mp7sihsuL",
	"B22i_VGhL9b",
	"B22hQKLAzup",
	"BvFQNBKnFtp",
	"B2mq1IuH6Jh",
	"B20fDT4gwgC",
	"B2h-k3CHKUf",
	"BxD3MFsFLKH",
	"B22i_YpB7ks",
	"BLUXLMQhNaN",
	"B20fWsUAjaC",
	"B2zvP--ha-j",
	"B21XOZpIMzS",
	"B22jDktlNLh",
	"B22jhDXnFBQ",
	"B21tkFolKLb",
	"B24KPXeFB-X",
	"B2mrQZrh_zS",
	"B20fWkxAoy_",
	"B22N_9wBa4n",
	"B22jhmFBcMz",
	"B23oEdJAfZO",
	"B24JzEMh3q7",
	"B2zkQ82g57o",
	"B24OOsDHA19",
	"B2J2sRdB6s_",
	"B20fXSjAkL3",
	"B2mriV9JccR",
	"B2moQZfAFGW",
	"B2mrcbAlm5l",
	"BzOrhpKAN16",
	"B2mgBLegHat",
	"BddLatxnrc2",
	"B2mX12KnO5L",
	"B8NN1JjH1rj",
	"B2mrstRna3p",
	"B20fXzIpYS7",
	"B2j3MGMg_vI",
	"B2mr8c7p59b",
	"B2mCFZoHlMz",
	"B6Kd5f2DApe",
	"B22jFKAg9pE",
	"B2msabdg95D",
	"B2mrpgiFVhl",
	"B2mrs4MFgcJ",
	"B2me6jXl6qE",
	"B23hHL1nD0K",
	"B2mr2-PFZjZ",
	"B2jcIp_oQcv",
	"B2mrxsNH92z",
	"B2HHetWnr6m",
	"B2mHQl7gZHl",
	"B2msMKGgqTl",
	"Bo7XhsZFzqu",
	"B2Z-v6PBjAo",
	"B2mr_UGnuN-",
	"B2msDKNAUpx",
	"B2fwgQaHymJ",
	"B20f3nGBcMy",
	"B22jIFRgxbK",
	"B20ffedhC6V",
	"B2mp-NMler4",
	"B2msZ8oACSX",
	"B2mrJvKlBZD",
	"Bzs9rpbBNjt",
	"B2zumIUhBpi",
	"B22jWTblJ_3",
	"B22jN_Xp8At",
	"B3AE8HZgkwf",
	"B2957csgaPg",
	"BrPKsNclZAH",
	"B0o6p6KDBYM",
	"B20gCqcpWA8",
	"B24J4PunQIA",
	"B20fot0hPwG",
	"B22jUPeBGph",
	"BlCYmbIj_hQ",
	"B24Jy6MHqqA",
	"B24J9z2FHCP",
	"B24KipJJR_9",
	"B2hJL_Zhvjy",
	"B2msBZ8AtWK",
	"B22jVYABm9e",
	"B23oJJ2iO4y",
	"B23oKdAoJXz",
	"B24J1HvnwLN",
	"B25T2jFgqTZ",
	"B25-rQEJmt_",
	"B262kriBx7X",
	"B2moZOJAqGj",
	"B2msdZohUBF",
	"B20ftpiJSeF",
	"B22jZ7WFzTK",
	"B23oNOFCXMs",
	"BpVtKIlg4pW",
	"B25l-7hHl8a",
	"B25-tFjgEwn",
	"B25mJnBpkhf",
	"B1_lqyZFt8j",
	"B2mr4R7g0ZT",
	"B2mr5r0hUzf",
	"BUElXrFBz35",
	"B2mr9T_AhGl",
	"B2iEbBCDl7b",
	"B2msIcjpcGf",
	"B2dVSroFGxq",
	"B2y-eYng2RX",
	"B22jQmKlm5p",
	"B23oTgppIyR",
	"B23oeV0Hrq7",
	"B24J6ImgN-_",
	"B2ckHgJH0Cb",
}
