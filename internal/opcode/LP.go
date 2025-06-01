package opcode

const (
	// LP_BEGIN_SOCKET
	LP_CheckPasswordResult        = 0
	LP_GuestIDLoginResult         = 1
	LP_AccountInfoResult          = 2
	LP_CheckUserLimitResult       = 3
	LP_SetAccountResult           = 4
	LP_ConfirmEULAResult          = 5
	LP_CheckPinCodeResult         = 6
	LP_UpdatePinCodeResult        = 7
	LP_ViewAllCharResult          = 8
	LP_SelectCharacterByVACResult = 9
	LP_WorldInformation           = 10
	LP_SelectWorldResult          = 11
	LP_SelectCharacterResult      = 12
	LP_CheckDuplicatedIDResult    = 13
	LP_CreateNewCharacterResult   = 14
	LP_DeleteCharacterResult      = 15
	LP_MigrateCommand             = 16
	LP_AliveReq                   = 17
	LP_AuthenCodeChanged          = 18
	LP_AuthenMessage              = 19
	LP_SecurityPacket             = 20
	LP_EnableSPWResult            = 21
	LP_DeleteCharacterOTPRequest  = 22
	LP_CheckCrcResult             = 23
	LP_LatestConnectedWorld       = 24
	LP_RecommendWorldMessage      = 25
	LP_CheckExtraCharInfoResult   = 26
	LP_CheckSPWResult             = 27
	// LP_END_SOCKET
	// LP_BEGIN_CHARACTERDATA
	LP_InventoryOperation            = 28
	LP_InventoryGrow                 = 29
	LP_StatChanged                   = 30
	LP_TemporaryStatSet              = 31
	LP_TemporaryStatReset            = 32
	LP_ForcedStatSet                 = 33
	LP_ForcedStatReset               = 34
	LP_ChangeSkillRecordResult       = 35
	LP_SkillUseResult                = 36
	LP_GivePopularityResult          = 37
	LP_Message                       = 38
	LP_SendOpenFullClientLink        = 39
	LP_MemoResult                    = 40
	LP_MapTransferResult             = 41
	LP_AntiMacroResult               = 42
	LP_InitialQuizStart              = 43
	LP_ClaimResult                   = 44
	LP_SetClaimSvrAvailableTime      = 45
	LP_ClaimSvrStatusChanged         = 46
	LP_SetTamingMobInfo              = 47
	LP_QuestClear                    = 48
	LP_EntrustedShopCheckResult      = 49
	LP_SkillLearnItemResult          = 50
	LP_SkillResetItemResult          = 51
	LP_GatherItemResult              = 52
	LP_SortItemResult                = 53
	LP_RemoteShopOpenResult          = 54
	LP_SueCharacterResult            = 55
	LP_MigrateToCashShopResult       = 56
	LP_TradeMoneyLimit               = 57
	LP_SetGender                     = 58
	LP_GuildBBS                      = 59
	LP_PetDeadMessage                = 60
	LP_CharacterInfo                 = 61
	LP_PartyResult                   = 62
	LP_ExpeditionRequest             = 63
	LP_ExpeditionNoti                = 64
	LP_FriendResult                  = 65
	LP_GuildRequest                  = 66
	LP_GuildResult                   = 67
	LP_AllianceResult                = 68
	LP_TownPortal                    = 69
	LP_OpenGate                      = 70
	LP_BroadcastMsg                  = 71
	LP_IncubatorResult               = 72
	LP_ShopScannerResult             = 73
	LP_ShopLinkResult                = 74
	LP_MarriageRequest               = 75
	LP_MarriageResult                = 76
	LP_WeddingGiftResult             = 77
	LP_MarriedPartnerMapTransfer     = 78
	LP_CashPetFoodResult             = 79
	LP_SetWeekEventMessage           = 80
	LP_SetPotionDiscountRate         = 81
	LP_BridleMobCatchFail            = 82
	LP_ImitatedNPCResult             = 83
	LP_ImitatedNPCData               = 84
	LP_LimitedNPCDisableInfo         = 85
	LP_MonsterBookSetCard            = 86
	LP_MonsterBookSetCover           = 87
	LP_HourChanged                   = 88
	LP_MiniMapOnOff                  = 89
	LP_ConsultAuthkeyUpdate          = 90
	LP_ClassCompetitionAuthkeyUpdate = 91
	LP_WebBoardAuthkeyUpdate         = 92
	LP_SessionValue                  = 93
	LP_PartyValue                    = 94
	LP_FieldSetVariable              = 95
	LP_BonusExpRateChanged           = 96
	LP_PotionDiscountRateChanged     = 97
	LP_FamilyChartResult             = 98
	LP_FamilyInfoResult              = 99
	LP_FamilyResult                  = 100
	LP_FamilyJoinRequest             = 101
	LP_FamilyJoinRequestResult       = 102
	LP_FamilyJoinAccepted            = 103
	LP_FamilyPrivilegeList           = 104
	LP_FamilyFamousPointIncResult    = 105
	LP_FamilyNotifyLoginOrLogout     = 106
	LP_FamilySetPrivilege            = 107
	LP_FamilySummonRequest           = 108
	LP_NotifyLevelUp                 = 109
	LP_NotifyWedding                 = 110
	LP_NotifyJobChange               = 111
	LP_IncRateChanged                = 112
	LP_MapleTVUseRes                 = 113
	LP_AvatarMegaphoneRes            = 114
	LP_AvatarMegaphoneUpdateMessage  = 115
	LP_AvatarMegaphoneClearMessage   = 116
	LP_CancelNameChangeResult        = 117
	LP_CancelTransferWorldResult     = 118
	LP_DestroyShopResult             = 119
	LP_FAKEGMNOTICE                  = 120
	LP_SuccessInUseGachaponBox       = 121
	LP_NewYearCardRes                = 122
	LP_RandomMorphRes                = 123
	LP_CancelNameChangeByOther       = 124
	LP_SetBuyEquipExt                = 125
	LP_SetPassenserRequest           = 126
	LP_ScriptProgressMessage         = 127
	LP_DataCRCCheckFailed            = 128
	LP_CakePieEventResult            = 129
	LP_UpdateGMBoard                 = 130
	LP_ShowSlotMessage               = 131
	LP_WildHunterInfo                = 132
	LP_AccountMoreInfo               = 133
	LP_FindFirend                    = 134
	LP_StageChange                   = 135
	LP_DragonBallBox                 = 136
	LP_AskUserWhetherUsePamsSong     = 137
	LP_TransferChannel               = 138
	LP_DisallowedDeliveryQuestList   = 139
	LP_MacroSysDataInit              = 140
	// LP_END_CHARACTERDATA
	// LP_BEGIN_STAGE
	LP_SetField    = 141
	LP_SetITC      = 142
	LP_SetCashShop = 143
	// LP_END_STAGE
	// LP_BEGIN_MAP
	LP_SetBackgroundEffect   = 144
	LP_SetMapObjectVisible   = 145
	LP_ClearBackgroundEffect = 146
	// LP_END_MAP
	// LP_BEGIN_FIELD
	LP_TransferFieldReqIgnored   = 147
	LP_TransferChannelReqIgnored = 148
	LP_FieldSpecificData         = 149
	LP_GroupMessage              = 150
	LP_Whisper                   = 151
	LP_CoupleMessage             = 152
	LP_MobSummonItemUseResult    = 153
	LP_FieldEffect               = 154
	LP_FieldObstacleOnOff        = 155
	LP_FieldObstacleOnOffStatus  = 156
	LP_FieldObstacleAllReset     = 157
	LP_BlowWeather               = 158
	LP_PlayJukeBox               = 159
	LP_AdminResult               = 160
	LP_Quiz                      = 161
	LP_Desc                      = 162
	LP_Clock                     = 163
	LP_CONTIMOVE                 = 164
	LP_CONTISTATE                = 165
	LP_SetQuestClear             = 166
	LP_SetQuestTime              = 167
	LP_Warn                      = 168
	LP_SetObjectState            = 169
	LP_DestroyClock              = 170
	LP_ShowArenaResult           = 171
	LP_StalkResult               = 172
	LP_MassacreIncGauge          = 173
	LP_MassacreResult            = 174
	LP_QuickslotMappedInit       = 175
	LP_FootHoldInfo              = 176
	LP_RequestFootHoldInfo       = 177
	LP_FieldKillCount            = 178
	// LP_BEGIN_USERPOOL
	LP_UserEnterField = 179
	LP_UserLeaveField = 180
	// LP_BEGIN_USERCOMMON
	LP_UserChat                    = 181
	LP_UserChatNLCPQ               = 182
	LP_UserADBoard                 = 183
	LP_UserMiniRoomBalloon         = 184
	LP_UserConsumeItemEffect       = 185
	LP_UserItemUpgradeEffect       = 186
	LP_UserItemHyperUpgradeEffect  = 187
	LP_UserItemOptionUpgradeEffect = 188
	LP_UserItemReleaseEffect       = 189
	LP_UserItemUnreleaseEffect     = 190
	LP_UserHitByUser               = 191
	LP_UserTeslaTriangle           = 192
	LP_UserFollowCharacter         = 193
	LP_UserShowPQReward            = 194
	LP_UserSetPhase                = 195
	LP_SetPortalUsable             = 196
	LP_ShowPamsSongResult          = 197
	// LP_BEGIN_PET
	LP_PetActivated         = 198
	LP_PetEvol              = 199
	LP_PetTransferField     = 200
	LP_PetMove              = 201
	LP_PetAction            = 202
	LP_PetNameChanged       = 203
	LP_PetLoadExceptionList = 204
	LP_PetActionCommand     = 205
	// LP_END_PET
	// LP_BEGIN_DRAGON
	LP_DragonEnterField = 206
	LP_DragonMove       = 207
	LP_DragonLeaveField = 208
	// LP_END_DRAGON
	// LP_END_USERCOMMON
	// LP_BEGIN_USERREMOTE
	LP_UserMove                     = 210
	LP_UserMeleeAttack              = 211
	LP_UserShootAttack              = 212
	LP_UserMagicAttack              = 213
	LP_UserBodyAttack               = 214
	LP_UserSkillPrepare             = 215
	LP_UserMovingShootAttackPrepare = 216
	LP_UserSkillCancel              = 217
	LP_UserHit                      = 218
	LP_UserEmotion                  = 219
	LP_UserSetActiveEffectItem      = 220
	LP_UserShowUpgradeTombEffect    = 221
	LP_UserSetActivePortableChair   = 222
	LP_UserAvatarModified           = 223
	LP_UserEffectRemote             = 224
	LP_UserTemporaryStatSet         = 225
	LP_UserTemporaryStatReset       = 226
	LP_UserHP                       = 227
	LP_UserGuildNameChanged         = 228
	LP_UserGuildMarkChanged         = 229
	LP_UserThrowGrenade             = 230
	// LP_END_USERREMOTE
	// LP_BEGIN_USERLOCAL
	LP_UserSitResult                = 231
	LP_UserEmotionLocal             = 232
	LP_UserEffectLocal              = 233
	LP_UserTeleport                 = 234
	LP_Premium                      = 235
	LP_MesoGive_Succeeded           = 236
	LP_MesoGive_Failed              = 237
	LP_Random_Mesobag_Succeed       = 238
	LP_Random_Mesobag_Failed        = 239
	LP_FieldFadeInOut               = 240
	LP_FieldFadeOutForce            = 241
	LP_UserQuestResult              = 242
	LP_NotifyHPDecByField           = 243
	LP_UserPetSkillChanged          = 244
	LP_UserBalloonMsg               = 245
	LP_PlayEventSound               = 246
	LP_PlayMinigameSound            = 247
	LP_UserMakerResult              = 248
	LP_UserOpenConsultBoard         = 249
	LP_UserOpenClassCompetitionPage = 250
	LP_UserOpenUI                   = 251
	LP_UserOpenUIWithOption         = 252
	LP_SetDirectionMode             = 253
	LP_SetStandAloneMode            = 254
	LP_UserHireTutor                = 255
	LP_UserTutorMsg                 = 256
	LP_IncCombo                     = 257
	LP_UserRandomEmotion            = 258
	LP_ResignQuestReturn            = 259
	LP_PassMateName                 = 260
	LP_SetRadioSchedule             = 261
	LP_UserOpenSkillGuide           = 262
	LP_UserNoticeMsg                = 263
	LP_UserChatMsg                  = 264
	LP_UserBuffzoneEffect           = 265
	LP_UserGoToCommoditySN          = 266
	LP_UserDamageMeter              = 267
	LP_UserTimeBombAttack           = 268
	LP_UserPassiveMove              = 269
	LP_UserFollowCharacterFailed    = 270
	LP_UserRequestVengeance         = 271
	LP_UserRequestExJablin          = 272
	LP_UserAskAPSPEvent             = 273
	LP_QuestGuideResult             = 274
	LP_UserDeliveryQuest            = 275
	LP_SkillCooltimeSet             = 276
	// LP_END_USERLOCAL
	// LP_END_USERPOOL
	// LP_BEGIN_SUMMONED
	LP_SummonedEnterField = 278
	LP_SummonedLeaveField = 279
	LP_SummonedMove       = 280
	LP_SummonedAttack     = 281
	LP_SummonedSkill      = 282
	LP_SummonedHit        = 283
	// LP_END_SUMMONED
	// LP_BEGIN_MOBPOOL
	LP_MobEnterField       = 284
	LP_MobLeaveField       = 285
	LP_MobChangeController = 286
	// LP_BEGIN_MOB
	LP_MobMove                    = 287
	LP_MobCtrlAck                 = 288
	LP_MobCtrlHint                = 289
	LP_MobStatSet                 = 290
	LP_MobStatReset               = 291
	LP_MobSuspendReset            = 292
	LP_MobAffected                = 293
	LP_MobDamaged                 = 294
	LP_MobSpecialEffectBySkill    = 295
	LP_MobHPChange                = 296
	LP_MobCrcKeyChanged           = 297
	LP_MobHPIndicator             = 298
	LP_MobCatchEffect             = 299
	LP_MobEffectByItem            = 300
	LP_MobSpeaking                = 301
	LP_MobChargeCount             = 302
	LP_MobSkillDelay              = 303
	LP_MobRequestResultEscortInfo = 304
	LP_MobEscortStopEndPermmision = 305
	LP_MobEscortStopSay           = 306
	LP_MobEscortReturnBefore      = 307
	LP_MobNextAttack              = 308
	LP_MobAttackedByMob           = 309
	// LP_END_MOB
	// LP_END_MOBPOOL
	// LP_BEGIN_NPCPOOL
	LP_NpcEnterField       = 311
	LP_NpcLeaveField       = 312
	LP_NpcChangeController = 313
	// LP_BEGIN_NPC
	LP_NpcMove              = 314
	LP_NpcUpdateLimitedInfo = 315
	LP_NpcSpecialAction     = 316
	// LP_END_NPC
	// LP_BEGIN_NPCTEMPLATE
	LP_NpcSetScript = 317
	// LP_END_NPCTEMPLATE
	// LP_END_NPCPOOL
	// LP_BEGIN_EMPLOYEEPOOL
	LP_EmployeeEnterField      = 319
	LP_EmployeeLeaveField      = 320
	LP_EmployeeMiniRoomBalloon = 321
	// LP_END_EMPLOYEEPOOL
	// LP_BEGIN_DROPPOOL
	LP_DropEnterField       = 322
	LP_DropReleaseAllFreeze = 323
	LP_DropLeaveField       = 324
	// LP_END_DROPPOOL
	// LP_BEGIN_MESSAGEBOXPOOL
	LP_CreateMessgaeBoxFailed = 325
	LP_MessageBoxEnterField   = 326
	LP_MessageBoxLeaveField   = 327
	// LP_END_MESSAGEBOXPOOL
	// LP_BEGIN_AFFECTEDAREAPOOL
	LP_AffectedAreaCreated = 328
	LP_AffectedAreaRemoved = 329
	// LP_END_AFFECTEDAREAPOOL
	// LP_BEGIN_TOWNPORTALPOOL
	LP_TownPortalCreated = 330
	LP_TownPortalRemoved = 331
	// LP_END_TOWNPORTALPOOL
	// LP_BEGIN_OPENGATEPOOL
	LP_OpenGateCreated = 332
	LP_OpenGateRemoved = 333
	// LP_END_OPENGATEPOOL
	// LP_BEGIN_REACTORPOOL
	LP_ReactorChangeState = 334
	LP_ReactorMove        = 335
	LP_ReactorEnterField  = 336
	LP_ReactorLeaveField  = 337
	// LP_END_REACTORPOOL
	// LP_BEGIN_ETCFIELDOBJ
	LP_SnowBallState          = 338
	LP_SnowBallHit            = 339
	LP_SnowBallMsg            = 340
	LP_SnowBallTouch          = 341
	LP_CoconutHit             = 342
	LP_CoconutScore           = 343
	LP_HealerMove             = 344
	LP_PulleyStateChange      = 345
	LP_MCarnivalEnter         = 346
	LP_MCarnivalPersonalCP    = 347
	LP_MCarnivalTeamCP        = 348
	LP_MCarnivalResultSuccess = 349
	LP_MCarnivalResultFail    = 350
	LP_MCarnivalDeath         = 351
	LP_MCarnivalMemberOut     = 352
	LP_MCarnivalGameResult    = 353
	LP_ArenaScore             = 354
	LP_BattlefieldEnter       = 355
	LP_BattlefieldScore       = 356
	LP_BattlefieldTeamChanged = 357
	LP_WitchtowerScore        = 358
	LP_HontaleTimer           = 359
	LP_ChaosZakumTimer        = 360
	LP_HontailTimer           = 361
	LP_ZakumTimer             = 362
	// LP_END_ETCFIELDOBJ
	// LP_BEGIN_SCRIPT
	LP_ScriptMessage = 363
	// LP_END_SCRIPT
	// LP_BEGIN_SHOP
	LP_OpenShopDlg = 364
	LP_ShopResult  = 365
	// LP_END_SHOP
	// LP_BEGIN_ADMINSHOP
	LP_AdminShopResult    = 366
	LP_AdminShopCommodity = 367
	// LP_END_ADMINSHOP
	LP_TrunkResult = 368
	// LP_BEGIN_STOREBANK
	LP_StoreBankGetAllResult = 369
	LP_StoreBankResult       = 370
	// LP_END_STOREBANK
	LP_RPSGame   = 371
	LP_Messenger = 372
	LP_MiniRoom  = 373
	// LP_BEGIN_TOURNAMENT
	LP_Tournament           = 374
	LP_TournamentMatchTable = 375
	LP_TournamentSetPrize   = 376
	LP_TournamentNoticeUEW  = 377
	LP_TournamentAvatarInfo = 378
	// LP_END_TOURNAMENT
	// LP_BEGIN_WEDDING
	LP_WeddingProgress   = 379
	LP_WeddingCremonyEnd = 380
	// LP_END_WEDDING
	LP_Parcel = 381
	// LP_END_FIELD
	// LP_BEGIN_CASHSHOP
	LP_CashShopChargeParamResult                = 382
	LP_CashShopQueryCashResult                  = 383
	LP_CashShopCashItemResult                   = 384
	LP_CashShopPurchaseExpChanged               = 385
	LP_CashShopGiftMateInfoResult               = 386
	LP_CashShopCheckDuplicatedIDResult          = 387
	LP_CashShopCheckNameChangePossibleResult    = 388
	LP_CashShopRegisterNewCharacterResult       = 389
	LP_CashShopCheckTransferWorldPossibleResult = 390
	LP_CashShopGachaponStampItemResult          = 391
	LP_CashShopCashItemGachaponResult           = 392
	LP_CashShopCashGachaponOpenResult           = 393
	LP_ChangeMaplePointResult                   = 394
	LP_CashShopOneADay                          = 395
	LP_CashShopNoticeFreeCashItem               = 396
	LP_CashShopMemberShopResult                 = 397
	// LP_END_CASHSHOP
	// LP_BEGIN_FUNCKEYMAPPED
	LP_FuncKeyMappedInit    = 398
	LP_PetConsumeItemInit   = 399
	LP_PetConsumeMPItemInit = 400
	// LP_END_FUNCKEYMAPPED
	LP_CheckSSN2OnCreateNewCharacterResult = 401
	LP_CheckSPWOnCreateNewCharacterResult  = 402
	LP_FirstSSNOnCreateNewCharacterResult  = 403
	// LP_BEGIN_MAPLETV
	LP_MapleTVUpdateMessage     = 405
	LP_MapleTVClearMessage      = 406
	LP_MapleTVSendMessageResult = 407
	LP_BroadSetFlashChangeEvent = 408
	// LP_END_MAPLETV
	// LP_BEGIN_ITC
	LP_ITCChargeParamResult = 410
	LP_ITCQueryCashResult   = 411
	LP_ITCNormalItemResult  = 412
	// LP_END_ITC
	// LP_BEGIN_CHARACTERSALE
	LP_CheckDuplicatedIDResultInCS  = 413
	LP_CreateNewCharacterResultInCS = 414
	LP_CreateNewCharacterFailInCS   = 415
	LP_CharacterSale                = 416
	// LP_END_CHARACTERSALE
	// LP_BEGIN_GOLDHAMMER
	LP_GoldHammere_s    = 417
	LP_GoldHammerResult = 418
	LP_GoldHammere_e    = 419
	// LP_END_GOLDHAMMER
	// LP_BEGIN_BATTLERECORD
	LP_BattleRecord_s            = 420
	LP_BattleRecordDotDamageInfo = 421
	LP_BattleRecordRequestResult = 422
	LP_BattleRecord_e            = 423
	// LP_END_BATTLERECORD
	// LP_BEGIN_ITEMUPGRADE
	LP_ItemUpgrade_s     = 424
	LP_ItemUpgradeResult = 425
	LP_ItemUpgradeFail   = 426
	LP_ItemUpgrade_e     = 427
	// LP_END_ITEMUPGRADE
	// LP_BEGIN_VEGA
	LP_Vega_s     = 428
	LP_VegaResult = 429
	LP_VegaFail   = 430
	LP_Vega_e     = 431
	// LP_END_VEGA
	LP_LogoutGift = 432
	LP_NO         = 433
)
