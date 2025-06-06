package opcode

const (
	CP_BEGIN_SOCKET                         = 0
	CP_CheckPassword                        = 1
	CP_GuestIDLogin                         = 2
	CP_AccountInfoRequest                   = 3
	CP_WorldInfoRequest                     = 4
	CP_SelectWorld                          = 5
	CP_CheckUserLimit                       = 6
	CP_ConfirmEULA                          = 7
	CP_SetGender                            = 8
	CP_CheckPinCode                         = 9
	CP_UpdatePinCode                        = 10
	CP_WorldRequest                         = 11
	CP_LogoutWorld                          = 12
	CP_ViewAllChar                          = 13
	CP_SelectCharacterByVAC                 = 14
	CP_VACFlagSet                           = 15
	CP_CheckNameChangePossible              = 16
	CP_RegisterNewCharacter                 = 17
	CP_CheckTransferWorldPossible           = 18
	CP_SelectCharacter                      = 19
	CP_MigrateIn                            = 20
	CP_CheckDuplicatedID                    = 21
	CP_CreateNewCharacter                   = 22
	CP_CreateNewCharacterInCS               = 23
	CP_DeleteCharacter                      = 24
	CP_AliveAck                             = 25
	CP_ExceptionLog                         = 26
	CP_SecurityPacket                       = 27
	CP_EnableSPWRequest                     = 28
	CP_CheckSPWRequest                      = 29
	CP_EnableSPWRequestByACV                = 30
	CP_CheckSPWRequestByACV                 = 31
	CP_CheckOTPRequest                      = 32
	CP_CheckDeleteCharacterOTP              = 33
	CP_CreateSecurityHandle                 = 34
	CP_SSOErrorLog                          = 35
	CP_ClientDumpLog                        = 36
	CP_CheckExtraCharInfo                   = 37
	CP_CreateNewCharacter_Ex                = 38
	CP_END_SOCKET                           = 39
	CP_BEGIN_USER                           = 40
	CP_UserTransferFieldRequest             = 41
	CP_UserTransferChannelRequest           = 42
	CP_UserMigrateToCashShopRequest         = 43
	CP_UserMove                             = 44
	CP_UserSitRequest                       = 45
	CP_UserPortableChairSitRequest          = 46
	CP_UserMeleeAttack                      = 47
	CP_UserShootAttack                      = 48
	CP_UserMagicAttack                      = 49
	CP_UserBodyAttack                       = 50
	CP_UserMovingShootAttackPrepare         = 51
	CP_UserHit                              = 52
	CP_UserAttackUser                       = 53
	CP_UserChat                             = 54
	CP_UserADBoardClose                     = 55
	CP_UserEmotion                          = 56
	CP_UserActivateEffectItem               = 57
	CP_UserUpgradeTombEffect                = 58
	CP_UserHP                               = 59
	CP_Premium                              = 60
	CP_UserBanMapByMob                      = 61
	CP_UserMonsterBookSetCover              = 62
	CP_UserSelectNpc                        = 63
	CP_UserRemoteShopOpenRequest            = 64
	CP_UserScriptMessageAnswer              = 65
	CP_UserShopRequest                      = 66
	CP_UserTrunkRequest                     = 67
	CP_UserEntrustedShopRequest             = 68
	CP_UserStoreBankRequest                 = 69
	CP_UserParcelRequest                    = 70
	CP_UserEffectLocal                      = 71
	CP_ShopScannerRequest                   = 72
	CP_ShopLinkRequest                      = 73
	CP_AdminShopRequest                     = 74
	CP_UserGatherItemRequest                = 75
	CP_UserSortItemRequest                  = 76
	CP_UserChangeSlotPositionRequest        = 77
	CP_UserStatChangeItemUseRequest         = 78
	CP_UserStatChangeItemCancelRequest      = 79
	CP_UserStatChangeByPortableChairRequest = 80
	CP_UserMobSummonItemUseRequest          = 81
	CP_UserPetFoodItemUseRequest            = 82
	CP_UserTamingMobFoodItemUseRequest      = 83
	CP_UserScriptItemUseRequest             = 84
	CP_UserConsumeCashItemUseRequest        = 85
	CP_UserDestroyPetItemRequest            = 86
	CP_UserBridleItemUseRequest             = 87
	CP_UserSkillLearnItemUseRequest         = 88
	CP_UserSkillResetItemUseRequest         = 89
	CP_UserShopScannerItemUseRequest        = 90
	CP_UserMapTransferItemUseRequest        = 91
	CP_UserPortalScrollUseRequest           = 92
	CP_UserUpgradeItemUseRequest            = 93
	CP_UserHyperUpgradeItemUseRequest       = 94
	CP_UserItemOptionUpgradeItemUseRequest  = 95
	CP_UserUIOpenItemUseRequest             = 96
	CP_UserItemReleaseRequest               = 97
	CP_UserAbilityUpRequest                 = 98
	CP_UserAbilityMassUpRequest             = 99
	CP_UserChangeStatRequest                = 100
	CP_UserChangeStatRequestByItemOption    = 101
	CP_UserSkillUpRequest                   = 102
	CP_UserSkillUseRequest                  = 103
	CP_UserSkillCancelRequest               = 104
	CP_UserSkillPrepareRequest              = 105
	CP_UserDropMoneyRequest                 = 106
	CP_UserGivePopularityRequest            = 107
	CP_UserPartyRequest                     = 108
	CP_UserCharacterInfoRequest             = 109
	CP_UserActivatePetRequest               = 110
	CP_UserTemporaryStatUpdateRequest       = 111
	CP_UserPortalScriptRequest              = 112
	CP_UserPortalTeleportRequest            = 113
	CP_UserMapTransferRequest               = 114
	CP_UserAntiMacroItemUseRequest          = 115
	CP_UserAntiMacroSkillUseRequest         = 116
	CP_UserAntiMacroQuestionResult          = 117
	CP_UserClaimRequest                     = 118
	CP_UserQuestRequest                     = 119
	CP_UserCalcDamageStatSetRequest         = 120
	CP_UserThrowGrenade                     = 121
	CP_UserMacroSysDataModified             = 122
	CP_UserSelectNpcItemUseRequest          = 123
	CP_UserLotteryItemUseRequest            = 124
	CP_UserItemMakeRequest                  = 125
	CP_UserSueCharacterRequest              = 126
	CP_UserUseGachaponBoxRequest            = 127
	CP_UserUseGachaponRemoteRequest         = 128
	CP_UserUseWaterOfLife                   = 129
	CP_UserRepairDurabilityAll              = 130
	CP_UserRepairDurability                 = 131
	CP_UserQuestRecordSetState              = 132
	CP_UserClientTimerEndRequest            = 133
	CP_UserFollowCharacterRequest           = 134
	CP_UserFollowCharacterWithdraw          = 135
	CP_UserSelectPQReward                   = 136
	CP_UserRequestPQReward                  = 137
	CP_SetPassenserResult                   = 138
	CP_BroadcastMsg                         = 139
	CP_GroupMessage                         = 140
	CP_Whisper                              = 141
	CP_CoupleMessage                        = 142
	CP_Messenger                            = 143
	CP_MiniRoom                             = 144
	CP_PartyRequest                         = 145
	CP_PartyResult                          = 146
	CP_ExpeditionRequest                    = 147
	CP_PartyAdverRequest                    = 148
	CP_GuildRequest                         = 149
	CP_GuildResult                          = 150
	CP_Admin                                = 151
	CP_Log                                  = 152
	CP_FriendRequest                        = 153
	CP_MemoRequest                          = 154
	CP_MemoFlagRequest                      = 155
	CP_EnterTownPortalRequest               = 156
	CP_EnterOpenGateRequest                 = 157
	CP_SlideRequest                         = 158
	CP_FuncKeyMappedModified                = 159
	CP_RPSGame                              = 160
	CP_MarriageRequest                      = 161
	CP_WeddingWishListRequest               = 162
	CP_WeddingProgress                      = 163
	CP_GuestBless                           = 164
	CP_BoobyTrapAlert                       = 165
	CP_StalkBegin                           = 166
	CP_AllianceRequest                      = 167
	CP_AllianceResult                       = 168
	CP_FamilyChartRequest                   = 169
	CP_FamilyInfoRequest                    = 170
	CP_FamilyRegisterJunior                 = 171
	CP_FamilyUnregisterJunior               = 172
	CP_FamilyUnregisterParent               = 173
	CP_FamilyJoinResult                     = 174
	CP_FamilyUsePrivilege                   = 175
	CP_FamilySetPrecept                     = 176
	CP_FamilySummonResult                   = 177
	CP_ChatBlockUserReq                     = 178
	CP_GuildBBS                             = 179
	CP_UserMigrateToITCRequest              = 180
	CP_UserExpUpItemUseRequest              = 181
	CP_UserTempExpUseRequest                = 182
	CP_NewYearCardRequest                   = 183
	CP_RandomMorphRequest                   = 184
	CP_CashItemGachaponRequest              = 185
	CP_CashGachaponOpenRequest              = 186
	CP_ChangeMaplePointRequest              = 187
	CP_TalkToTutor                          = 188
	CP_RequestIncCombo                      = 189
	CP_MobCrcKeyChangedReply                = 190
	CP_RequestSessionValue                  = 191
	CP_UpdateGMBoard                        = 192
	CP_AccountMoreInfo                      = 193
	CP_FindFriend                           = 194
	CP_AcceptAPSPEvent                      = 195
	CP_UserDragonBallBoxRequest             = 196
	CP_UserDragonBallSummonRequest          = 197
	CP_BEGIN_PET                            = 198
	CP_PetMove                              = 199
	CP_PetAction                            = 200
	CP_PetInteractionRequest                = 201
	CP_PetDropPickUpRequest                 = 202
	CP_PetStatChangeItemUseRequest          = 203
	CP_PetUpdateExceptionListRequest        = 204
	CP_END_PET                              = 205
	CP_BEGIN_SUMMONED                       = 206
	CP_SummonedMove                         = 207
	CP_SummonedAttack                       = 208
	CP_SummonedHit                          = 209
	CP_SummonedSkill                        = 210
	CP_Remove                               = 211
	CP_END_SUMMONED                         = 212
	CP_BEGIN_DRAGON                         = 213
	CP_DragonMove                           = 214
	CP_END_DRAGON                           = 215
	CP_QuickslotKeyMappedModified           = 216
	CP_PassiveskillInfoUpdate               = 217
	CP_UpdateScreenSetting                  = 218
	CP_UserAttackUser_Specific              = 219
	CP_UserPamsSongUseRequest               = 220
	CP_QuestGuideRequest                    = 221
	CP_UserRepeatEffectRemove               = 222
	CP_END_USER                             = 223
	CP_BEGIN_FIELD                          = 224
	CP_BEGIN_LIFEPOOL                       = 225
	CP_BEGIN_MOB                            = 226
	CP_MobMove                              = 227
	CP_MobApplyCtrl                         = 228
	CP_MobDropPickUpRequest                 = 229
	CP_MobHitByObstacle                     = 230
	CP_MobHitByMob                          = 231
	CP_MobSelfDestruct                      = 232
	CP_MobAttackMob                         = 233
	CP_MobSkillDelayEnd                     = 234
	CP_MobTimeBombEnd                       = 235
	CP_MobEscortCollision                   = 236
	CP_MobRequestEscortInfo                 = 237
	CP_MobEscortStopEndRequest              = 238
	CP_END_MOB                              = 239
	CP_BEGIN_NPC                            = 240
	CP_NpcMove                              = 241
	CP_NpcSpecialAction                     = 242
	CP_END_NPC                              = 243
	CP_END_LIFEPOOL                         = 244
	CP_BEGIN_DROPPOOL                       = 245
	CP_DropPickUpRequest                    = 246
	CP_END_DROPPOOL                         = 247
	CP_BEGIN_REACTORPOOL                    = 248
	CP_ReactorHit                           = 249
	CP_ReactorTouch                         = 250
	CP_RequireFieldObstacleStatus           = 251
	CP_END_REACTORPOOL                      = 252
	CP_BEGIN_EVENT_FIELD                    = 253
	CP_EventStart                           = 254
	CP_SnowBallHit                          = 255
	CP_SnowBallTouch                        = 256
	CP_CoconutHit                           = 257
	CP_TournamentMatchTable                 = 258
	CP_PulleyHit                            = 259
	CP_END_EVENT_FIELD                      = 260
	CP_BEGIN_MONSTER_CARNIVAL_FIELD         = 261
	CP_MCarnivalRequest                     = 262
	CP_END_MONSTER_CARNIVAL_FIELD           = 263
	CP_CONTISTATE                           = 264
	CP_BEGIN_PARTY_MATCH                    = 265
	CP_INVITE_PARTY_MATCH                   = 266
	CP_CANCEL_INVITE_PARTY_MATCH            = 267
	CP_END_PARTY_MATCH                      = 268
	CP_RequestFootHoldInfo                  = 269
	CP_FootHoldInfo                         = 270
	CP_END_FIELD                            = 271
	CP_BEGIN_CASHSHOP                       = 272
	CP_CashShopChargeParamRequest           = 273
	CP_CashShopQueryCashRequest             = 274
	CP_CashShopCashItemRequest              = 275
	CP_CashShopCheckCouponRequest           = 276
	CP_CashShopGiftMateInfoRequest          = 277
	CP_END_CASHSHOP                         = 278
	CP_CheckSSN2OnCreateNewCharacter        = 279
	CP_CheckSPWOnCreateNewCharacter         = 280
	CP_FirstSSNOnCreateNewCharacter         = 281
	CP_BEGIN_RAISE                          = 282
	CP_RaiseRefesh                          = 283
	CP_RaiseUIState                         = 284
	CP_RaiseIncExp                          = 285
	CP_RaiseAddPiece                        = 286
	CP_END_RAISE                            = 287
	CP_SendMateMail                         = 288
	CP_RequestGuildBoardAuthKey             = 289
	CP_RequestConsultAuthKey                = 290
	CP_RequestClassCompetitionAuthKey       = 291
	CP_RequestWebBoardAuthKey               = 292
	CP_BEGIN_ITEMUPGRADE                    = 293
	CP_GoldHammerRequest                    = 294
	CP_GoldHammerComplete                   = 295
	CP_ItemUpgradeComplete                  = 296
	CP_END_ITEMUPGRADE                      = 297
	CP_BEGIN_BATTLERECORD                   = 298
	CP_BATTLERECORD_ONOFF_REQUEST           = 299
	CP_END_BATTLERECORD                     = 300
	CP_BEGIN_MAPLETV                        = 301
	CP_MapleTVSendMessageRequest            = 302
	CP_MapleTVUpdateViewCount               = 303
	CP_END_MAPLETV                          = 304
	CP_BEGIN_ITC                            = 305
	CP_ITCChargeParamRequest                = 306
	CP_ITCQueryCashRequest                  = 307
	CP_ITCItemRequest                       = 308
	CP_END_ITC                              = 309
	CP_BEGIN_CHARACTERSALE                  = 310
	CP_CheckDuplicatedIDInCS                = 311
	CP_END_CHARACTERSALE                    = 312
	CP_LogoutGiftSelect                     = 313
	CP_NO                                   = 314
)
