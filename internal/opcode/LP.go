package opcode

const (
	// LP_BEGIN_SOCKET
	LP_CheckPasswordResult        = 0x0
	LP_GuestIDLoginResult         = 0x1
	LP_AccountInfoResult          = 0x2
	LP_CheckUserLimitResult       = 0x3
	LP_SetAccountResult           = 0x4
	LP_ConfirmEULAResult          = 0x5
	LP_CheckPinCodeResult         = 0x6
	LP_UpdatePinCodeResult        = 0x7
	LP_ViewAllCharResult          = 0x8
	LP_SelectCharacterByVACResult = 0x9
	LP_WorldInformation           = 0xA
	LP_SelectWorldResult          = 0xB
	LP_SelectCharacterResult      = 0xC
	LP_CheckDuplicatedIDResult    = 0xD
	LP_CreateNewCharacterResult   = 0xE
	LP_DeleteCharacterResult      = 0xF
	LP_MigrateCommand             = 0x10
	LP_AliveReq                   = 0x11
	LP_AuthenCodeChanged          = 0x12
	LP_AuthenMessage              = 0x13
	LP_SecurityPacket             = 0x14
	LP_EnableSPWResult            = 0x15
	LP_DeleteCharacterOTPRequest  = 0x16
	LP_CheckCrcResult             = 0x17
	LP_LatestConnectedWorld       = 0x18
	LP_RecommendWorldMessage      = 0x19
	LP_CheckExtraCharInfoResult   = 0x1A
	LP_CheckSPWResult             = 0x1B
	// LP_END_SOCKET
	// LP_BEGIN_CHARACTERDATA
	LP_InventoryOperation            = 0x1C
	LP_InventoryGrow                 = 0x1D
	LP_StatChanged                   = 0x1E
	LP_TemporaryStatSet              = 0x1F
	LP_TemporaryStatReset            = 0x20
	LP_ForcedStatSet                 = 0x21
	LP_ForcedStatReset               = 0x22
	LP_ChangeSkillRecordResult       = 0x23
	LP_SkillUseResult                = 0x24
	LP_GivePopularityResult          = 0x25
	LP_Message                       = 0x26
	LP_SendOpenFullClientLink        = 0x27
	LP_MemoResult                    = 0x28
	LP_MapTransferResult             = 0x29
	LP_AntiMacroResult               = 0x2A
	LP_InitialQuizStart              = 0x2B
	LP_ClaimResult                   = 0x2C
	LP_SetClaimSvrAvailableTime      = 0x2D
	LP_ClaimSvrStatusChanged         = 0x2E
	LP_SetTamingMobInfo              = 0x2F
	LP_QuestClear                    = 0x30
	LP_EntrustedShopCheckResult      = 0x31
	LP_SkillLearnItemResult          = 0x32
	LP_SkillResetItemResult          = 0x33
	LP_GatherItemResult              = 0x34
	LP_SortItemResult                = 0x35
	LP_RemoteShopOpenResult          = 0x36
	LP_SueCharacterResult            = 0x37
	LP_MigrateToCashShopResult       = 0x38
	LP_TradeMoneyLimit               = 0x39
	LP_SetGender                     = 0x3A
	LP_GuildBBS                      = 0x3B
	LP_PetDeadMessage                = 0x3C
	LP_CharacterInfo                 = 0x3D
	LP_PartyResult                   = 0x3E
	LP_ExpeditionRequest             = 0x3F
	LP_ExpeditionNoti                = 0x40
	LP_FriendResult                  = 0x41
	LP_GuildRequest                  = 0x42
	LP_GuildResult                   = 0x43
	LP_AllianceResult                = 0x44
	LP_TownPortal                    = 0x45
	LP_OpenGate                      = 0x46
	LP_BroadcastMsg                  = 0x47
	LP_IncubatorResult               = 0x48
	LP_ShopScannerResult             = 0x49
	LP_ShopLinkResult                = 0x4A
	LP_MarriageRequest               = 0x4B
	LP_MarriageResult                = 0x4C
	LP_WeddingGiftResult             = 0x4D
	LP_MarriedPartnerMapTransfer     = 0x4E
	LP_CashPetFoodResult             = 0x4F
	LP_SetWeekEventMessage           = 0x50
	LP_SetPotionDiscountRate         = 0x51
	LP_BridleMobCatchFail            = 0x52
	LP_ImitatedNPCResult             = 0x53
	LP_ImitatedNPCData               = 0x54
	LP_LimitedNPCDisableInfo         = 0x55
	LP_MonsterBookSetCard            = 0x56
	LP_MonsterBookSetCover           = 0x57
	LP_HourChanged                   = 0x58
	LP_MiniMapOnOff                  = 0x59
	LP_ConsultAuthkeyUpdate          = 0x5A
	LP_ClassCompetitionAuthkeyUpdate = 0x5B
	LP_WebBoardAuthkeyUpdate         = 0x5C
	LP_SessionValue                  = 0x5D
	LP_PartyValue                    = 0x5E
	LP_FieldSetVariable              = 0x5F
	LP_BonusExpRateChanged           = 0x60
	LP_PotionDiscountRateChanged     = 0x61
	LP_FamilyChartResult             = 0x62
	LP_FamilyInfoResult              = 0x63
	LP_FamilyResult                  = 0x64
	LP_FamilyJoinRequest             = 0x65
	LP_FamilyJoinRequestResult       = 0x66
	LP_FamilyJoinAccepted            = 0x67
	LP_FamilyPrivilegeList           = 0x68
	LP_FamilyFamousPointIncResult    = 0x69
	LP_FamilyNotifyLoginOrLogout     = 0x6A
	LP_FamilySetPrivilege            = 0x6B
	LP_FamilySummonRequest           = 0x6C
	LP_NotifyLevelUp                 = 0x6D
	LP_NotifyWedding                 = 0x6E
	LP_NotifyJobChange               = 0x6F
	LP_IncRateChanged                = 0x70
	LP_MapleTVUseRes                 = 0x71
	LP_AvatarMegaphoneRes            = 0x72
	LP_AvatarMegaphoneUpdateMessage  = 0x73
	LP_AvatarMegaphoneClearMessage   = 0x74
	LP_CancelNameChangeResult        = 0x75
	LP_CancelTransferWorldResult     = 0x76
	LP_DestroyShopResult             = 0x77
	LP_FAKEGMNOTICE                  = 0x78
	LP_SuccessInUseGachaponBox       = 0x79
	LP_NewYearCardRes                = 0x7A
	LP_RandomMorphRes                = 0x7B
	LP_CancelNameChangeByOther       = 0x7C
	LP_SetBuyEquipExt                = 0x7D
	LP_SetPassenserRequest           = 0x7E
	LP_ScriptProgressMessage         = 0x7F
	LP_DataCRCCheckFailed            = 0x80
	LP_CakePieEventResult            = 0x81
	LP_UpdateGMBoard                 = 0x82
	LP_ShowSlotMessage               = 0x83
	LP_WildHunterInfo                = 0x84
	LP_AccountMoreInfo               = 0x85
	LP_FindFirend                    = 0x86
	LP_StageChange                   = 0x87
	LP_DragonBallBox                 = 0x88
	LP_AskUserWhetherUsePamsSong     = 0x89
	LP_TransferChannel               = 0x8A
	LP_DisallowedDeliveryQuestList   = 0x8B
	LP_MacroSysDataInit              = 0x8C
	// LP_END_CHARACTERDATA
	// LP_BEGIN_STAGE
	LP_SetField    = 0x8D
	LP_SetITC      = 0x8E
	LP_SetCashShop = 0x8F
	// LP_END_STAGE
	// LP_BEGIN_MAP
	LP_SetBackgroundEffect   = 0x90
	LP_SetMapObjectVisible   = 0x91
	LP_ClearBackgroundEffect = 0x92
	// LP_END_MAP
	// LP_BEGIN_FIELD
	LP_TransferFieldReqIgnored   = 0x93
	LP_TransferChannelReqIgnored = 0x94
	LP_FieldSpecificData         = 0x95
	LP_GroupMessage              = 0x96
	LP_Whisper                   = 0x97
	LP_CoupleMessage             = 0x98
	LP_MobSummonItemUseResult    = 0x99
	LP_FieldEffect               = 0x9A
	LP_FieldObstacleOnOff        = 0x9B
	LP_FieldObstacleOnOffStatus  = 0x9C
	LP_FieldObstacleAllReset     = 0x9D
	LP_BlowWeather               = 0x9E
	LP_PlayJukeBox               = 0x9F
	LP_AdminResult               = 0xA0
	LP_Quiz                      = 0xA1
	LP_Desc                      = 0xA2
	LP_Clock                     = 0xA3
	LP_CONTIMOVE                 = 0xA4
	LP_CONTISTATE                = 0xA5
	LP_SetQuestClear             = 0xA6
	LP_SetQuestTime              = 0xA7
	LP_Warn                      = 0xA8
	LP_SetObjectState            = 0xA9
	LP_DestroyClock              = 0xAA
	LP_ShowArenaResult           = 0xAB
	LP_StalkResult               = 0xAC
	LP_MassacreIncGauge          = 0xAD
	LP_MassacreResult            = 0xAE
	LP_QuickslotMappedInit       = 0xAF
	LP_FootHoldInfo              = 0xB0
	LP_RequestFootHoldInfo       = 0xB1
	LP_FieldKillCount            = 0xB2
	// LP_BEGIN_USERPOOL
	LP_UserEnterField = 0xB3
	LP_UserLeaveField = 0xB4
	// LP_BEGIN_USERCOMMON
	LP_UserChat                    = 0xB5
	LP_UserChatNLCPQ               = 0xB6
	LP_UserADBoard                 = 0xB7
	LP_UserMiniRoomBalloon         = 0xB8
	LP_UserConsumeItemEffect       = 0xB9
	LP_UserItemUpgradeEffect       = 0xBA
	LP_UserItemHyperUpgradeEffect  = 0xBB
	LP_UserItemOptionUpgradeEffect = 0xBC
	LP_UserItemReleaseEffect       = 0xBD
	LP_UserItemUnreleaseEffect     = 0xBE
	LP_UserHitByUser               = 0xBF
	LP_UserTeslaTriangle           = 0xC0
	LP_UserFollowCharacter         = 0xC1
	LP_UserShowPQReward            = 0xC2
	LP_UserSetPhase                = 0xC3
	LP_SetPortalUsable             = 0xC4
	LP_ShowPamsSongResult          = 0xC5
	// LP_BEGIN_PET
	LP_PetActivated         = 0xC6
	LP_PetEvol              = 0xC7
	LP_PetTransferField     = 0xC8
	LP_PetMove              = 0xC9
	LP_PetAction            = 0xCA
	LP_PetNameChanged       = 0xCB
	LP_PetLoadExceptionList = 0xCC
	LP_PetActionCommand     = 0xCD
	// LP_END_PET
	// LP_BEGIN_DRAGON
	LP_DragonEnterField = 0xCE
	LP_DragonMove       = 0xCF
	LP_DragonLeaveField = 0xD0
	// LP_END_DRAGON
	// LP_END_USERCOMMON
	// LP_BEGIN_USERREMOTE
	LP_UserMove                     = 0xD2
	LP_UserMeleeAttack              = 0xD3
	LP_UserShootAttack              = 0xD4
	LP_UserMagicAttack              = 0xD5
	LP_UserBodyAttack               = 0xD6
	LP_UserSkillPrepare             = 0xD7
	LP_UserMovingShootAttackPrepare = 0xD8
	LP_UserSkillCancel              = 0xD9
	LP_UserHit                      = 0xDA
	LP_UserEmotion                  = 0xDB
	LP_UserSetActiveEffectItem      = 0xDC
	LP_UserShowUpgradeTombEffect    = 0xDD
	LP_UserSetActivePortableChair   = 0xDE
	LP_UserAvatarModified           = 0xDF
	LP_UserEffectRemote             = 0xE0
	LP_UserTemporaryStatSet         = 0xE1
	LP_UserTemporaryStatReset       = 0xE2
	LP_UserHP                       = 0xE3
	LP_UserGuildNameChanged         = 0xE4
	LP_UserGuildMarkChanged         = 0xE5
	LP_UserThrowGrenade             = 0xE6
	// LP_END_USERREMOTE
	// LP_BEGIN_USERLOCAL
	LP_UserSitResult                = 0xE7
	LP_UserEmotionLocal             = 0xE8
	LP_UserEffectLocal              = 0xE9
	LP_UserTeleport                 = 0xEA
	LP_Premium                      = 0xEB
	LP_MesoGive_Succeeded           = 0xEC
	LP_MesoGive_Failed              = 0xED
	LP_Random_Mesobag_Succeed       = 0xEE
	LP_Random_Mesobag_Failed        = 0xEF
	LP_FieldFadeInOut               = 0xF0
	LP_FieldFadeOutForce            = 0xF1
	LP_UserQuestResult              = 0xF2
	LP_NotifyHPDecByField           = 0xF3
	LP_UserPetSkillChanged          = 0xF4
	LP_UserBalloonMsg               = 0xF5
	LP_PlayEventSound               = 0xF6
	LP_PlayMinigameSound            = 0xF7
	LP_UserMakerResult              = 0xF8
	LP_UserOpenConsultBoard         = 0xF9
	LP_UserOpenClassCompetitionPage = 0xFA
	LP_UserOpenUI                   = 0xFB
	LP_UserOpenUIWithOption         = 0xFC
	LP_SetDirectionMode             = 0xFD
	LP_SetStandAloneMode            = 0xFE
	LP_UserHireTutor                = 0xFF
	LP_UserTutorMsg                 = 0x100
	LP_IncCombo                     = 0x101
	LP_UserRandomEmotion            = 0x102
	LP_ResignQuestReturn            = 0x103
	LP_PassMateName                 = 0x104
	LP_SetRadioSchedule             = 0x105
	LP_UserOpenSkillGuide           = 0x106
	LP_UserNoticeMsg                = 0x107
	LP_UserChatMsg                  = 0x108
	LP_UserBuffzoneEffect           = 0x109
	LP_UserGoToCommoditySN          = 0x10A
	LP_UserDamageMeter              = 0x10B
	LP_UserTimeBombAttack           = 0x10C
	LP_UserPassiveMove              = 0x10D
	LP_UserFollowCharacterFailed    = 0x10E
	LP_UserRequestVengeance         = 0x10F
	LP_UserRequestExJablin          = 0x110
	LP_UserAskAPSPEvent             = 0x111
	LP_QuestGuideResult             = 0x112
	LP_UserDeliveryQuest            = 0x113
	LP_SkillCooltimeSet             = 0x114
	// LP_END_USERLOCAL
	// LP_END_USERPOOL
	// LP_BEGIN_SUMMONED
	LP_SummonedEnterField = 0x116
	LP_SummonedLeaveField = 0x117
	LP_SummonedMove       = 0x118
	LP_SummonedAttack     = 0x119
	LP_SummonedSkill      = 0x11A
	LP_SummonedHit        = 0x11B
	// LP_END_SUMMONED
	// LP_BEGIN_MOBPOOL
	LP_MobEnterField       = 0x11C
	LP_MobLeaveField       = 0x11D
	LP_MobChangeController = 0x11E
	// LP_BEGIN_MOB
	LP_MobMove                    = 0x11F
	LP_MobCtrlAck                 = 0x120
	LP_MobCtrlHint                = 0x121
	LP_MobStatSet                 = 0x122
	LP_MobStatReset               = 0x123
	LP_MobSuspendReset            = 0x124
	LP_MobAffected                = 0x125
	LP_MobDamaged                 = 0x126
	LP_MobSpecialEffectBySkill    = 0x127
	LP_MobHPChange                = 0x128
	LP_MobCrcKeyChanged           = 0x129
	LP_MobHPIndicator             = 0x12A
	LP_MobCatchEffect             = 0x12B
	LP_MobEffectByItem            = 0x12C
	LP_MobSpeaking                = 0x12D
	LP_MobChargeCount             = 0x12E
	LP_MobSkillDelay              = 0x12F
	LP_MobRequestResultEscortInfo = 0x130
	LP_MobEscortStopEndPermmision = 0x131
	LP_MobEscortStopSay           = 0x132
	LP_MobEscortReturnBefore      = 0x133
	LP_MobNextAttack              = 0x134
	LP_MobAttackedByMob           = 0x135
	// LP_END_MOB
	// LP_END_MOBPOOL
	// LP_BEGIN_NPCPOOL
	LP_NpcEnterField       = 0x137
	LP_NpcLeaveField       = 0x138
	LP_NpcChangeController = 0x139
	// LP_BEGIN_NPC
	LP_NpcMove              = 0x13A
	LP_NpcUpdateLimitedInfo = 0x13B
	LP_NpcSpecialAction     = 0x13C
	// LP_END_NPC
	// LP_BEGIN_NPCTEMPLATE
	LP_NpcSetScript = 0x13D
	// LP_END_NPCTEMPLATE
	// LP_END_NPCPOOL
	// LP_BEGIN_EMPLOYEEPOOL
	LP_EmployeeEnterField      = 0x13F
	LP_EmployeeLeaveField      = 0x140
	LP_EmployeeMiniRoomBalloon = 0x141
	// LP_END_EMPLOYEEPOOL
	// LP_BEGIN_DROPPOOL
	LP_DropEnterField       = 0x142
	LP_DropReleaseAllFreeze = 0x143
	LP_DropLeaveField       = 0x144
	// LP_END_DROPPOOL
	// LP_BEGIN_MESSAGEBOXPOOL
	LP_CreateMessgaeBoxFailed = 0x145
	LP_MessageBoxEnterField   = 0x146
	LP_MessageBoxLeaveField   = 0x147
	// LP_END_MESSAGEBOXPOOL
	// LP_BEGIN_AFFECTEDAREAPOOL
	LP_AffectedAreaCreated = 0x148
	LP_AffectedAreaRemoved = 0x149
	// LP_END_AFFECTEDAREAPOOL
	// LP_BEGIN_TOWNPORTALPOOL
	LP_TownPortalCreated = 0x14A
	LP_TownPortalRemoved = 0x14B
	// LP_END_TOWNPORTALPOOL
	// LP_BEGIN_OPENGATEPOOL
	LP_OpenGateCreated = 0x14C
	LP_OpenGateRemoved = 0x14D
	// LP_END_OPENGATEPOOL
	// LP_BEGIN_REACTORPOOL
	LP_ReactorChangeState = 0x14E
	LP_ReactorMove        = 0x14F
	LP_ReactorEnterField  = 0x150
	LP_ReactorLeaveField  = 0x151
	// LP_END_REACTORPOOL
	// LP_BEGIN_ETCFIELDOBJ
	LP_SnowBallState          = 0x152
	LP_SnowBallHit            = 0x153
	LP_SnowBallMsg            = 0x154
	LP_SnowBallTouch          = 0x155
	LP_CoconutHit             = 0x156
	LP_CoconutScore           = 0x157
	LP_HealerMove             = 0x158
	LP_PulleyStateChange      = 0x159
	LP_MCarnivalEnter         = 0x15A
	LP_MCarnivalPersonalCP    = 0x15B
	LP_MCarnivalTeamCP        = 0x15C
	LP_MCarnivalResultSuccess = 0x15D
	LP_MCarnivalResultFail    = 0x15E
	LP_MCarnivalDeath         = 0x15F
	LP_MCarnivalMemberOut     = 0x160
	LP_MCarnivalGameResult    = 0x161
	LP_ArenaScore             = 0x162
	LP_BattlefieldEnter       = 0x163
	LP_BattlefieldScore       = 0x164
	LP_BattlefieldTeamChanged = 0x165
	LP_WitchtowerScore        = 0x166
	LP_HontaleTimer           = 0x167
	LP_ChaosZakumTimer        = 0x168
	LP_HontailTimer           = 0x169
	LP_ZakumTimer             = 0x16A
	// LP_END_ETCFIELDOBJ
	// LP_BEGIN_SCRIPT
	LP_ScriptMessage = 0x16B
	// LP_END_SCRIPT
	// LP_BEGIN_SHOP
	LP_OpenShopDlg = 0x16C
	LP_ShopResult  = 0x16D
	// LP_END_SHOP
	// LP_BEGIN_ADMINSHOP
	LP_AdminShopResult    = 0x16E
	LP_AdminShopCommodity = 0x16F
	// LP_END_ADMINSHOP
	LP_TrunkResult = 0x170
	// LP_BEGIN_STOREBANK
	LP_StoreBankGetAllResult = 0x171
	LP_StoreBankResult       = 0x172
	// LP_END_STOREBANK
	LP_RPSGame   = 0x173
	LP_Messenger = 0x174
	LP_MiniRoom  = 0x175
	// LP_BEGIN_TOURNAMENT
	LP_Tournament           = 0x176
	LP_TournamentMatchTable = 0x177
	LP_TournamentSetPrize   = 0x178
	LP_TournamentNoticeUEW  = 0x179
	LP_TournamentAvatarInfo = 0x17A
	// LP_END_TOURNAMENT
	// LP_BEGIN_WEDDING
	LP_WeddingProgress   = 0x17B
	LP_WeddingCremonyEnd = 0x17C
	// LP_END_WEDDING
	LP_Parcel = 0x17D
	// LP_END_FIELD
	// LP_BEGIN_CASHSHOP
	LP_CashShopChargeParamResult                = 0x17E
	LP_CashShopQueryCashResult                  = 0x17F
	LP_CashShopCashItemResult                   = 0x180
	LP_CashShopPurchaseExpChanged               = 0x181
	LP_CashShopGiftMateInfoResult               = 0x182
	LP_CashShopCheckDuplicatedIDResult          = 0x183
	LP_CashShopCheckNameChangePossibleResult    = 0x184
	LP_CashShopRegisterNewCharacterResult       = 0x185
	LP_CashShopCheckTransferWorldPossibleResult = 0x186
	LP_CashShopGachaponStampItemResult          = 0x187
	LP_CashShopCashItemGachaponResult           = 0x188
	LP_CashShopCashGachaponOpenResult           = 0x189
	LP_ChangeMaplePointResult                   = 0x18A
	LP_CashShopOneADay                          = 0x18B
	LP_CashShopNoticeFreeCashItem               = 0x18C
	LP_CashShopMemberShopResult                 = 0x18D
	// LP_END_CASHSHOP
	// LP_BEGIN_FUNCKEYMAPPED
	LP_FuncKeyMappedInit    = 0x18E
	LP_PetConsumeItemInit   = 0x18F
	LP_PetConsumeMPItemInit = 0x190
	// LP_END_FUNCKEYMAPPED
	LP_CheckSSN2OnCreateNewCharacterResult = 0x191
	LP_CheckSPWOnCreateNewCharacterResult  = 0x192
	LP_FirstSSNOnCreateNewCharacterResult  = 0x193
	// LP_BEGIN_MAPLETV
	LP_MapleTVUpdateMessage     = 0x195
	LP_MapleTVClearMessage      = 0x196
	LP_MapleTVSendMessageResult = 0x197
	LP_BroadSetFlashChangeEvent = 0x198
	// LP_END_MAPLETV
	// LP_BEGIN_ITC
	LP_ITCChargeParamResult = 0x19A
	LP_ITCQueryCashResult   = 0x19B
	LP_ITCNormalItemResult  = 0x19C
	// LP_END_ITC
	// LP_BEGIN_CHARACTERSALE
	LP_CheckDuplicatedIDResultInCS  = 0x19D
	LP_CreateNewCharacterResultInCS = 0x19E
	LP_CreateNewCharacterFailInCS   = 0x19F
	LP_CharacterSale                = 0x1A0
	// LP_END_CHARACTERSALE
	// LP_BEGIN_GOLDHAMMER
	LP_GoldHammere_s    = 0x1A1
	LP_GoldHammerResult = 0x1A2
	LP_GoldHammere_e    = 0x1A3
	// LP_END_GOLDHAMMER
	// LP_BEGIN_BATTLERECORD
	LP_BattleRecord_s            = 0x1A4
	LP_BattleRecordDotDamageInfo = 0x1A5
	LP_BattleRecordRequestResult = 0x1A6
	LP_BattleRecord_e            = 0x1A7
	// LP_END_BATTLERECORD
	// LP_BEGIN_ITEMUPGRADE
	LP_ItemUpgrade_s     = 0x1A8
	LP_ItemUpgradeResult = 0x1A9
	LP_ItemUpgradeFail   = 0x1AA
	LP_ItemUpgrade_e     = 0x1AB
	// LP_END_ITEMUPGRADE
	// LP_BEGIN_VEGA
	LP_Vega_s     = 0x1AC
	LP_VegaResult = 0x1AD
	LP_VegaFail   = 0x1AE
	LP_Vega_e     = 0x1AF
	// LP_END_VEGA
	LP_LogoutGift = 0x1B0
	LP_NO         = 0x1B1
)
