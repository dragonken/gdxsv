// Code generated by "stringer -type=CmdID"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[lbsLineCheck-24577]
	_ = x[lbsLogout-24578]
	_ = x[lbsShutDown-24579]
	_ = x[lbsVSUserLost-24580]
	_ = x[lbsSendMail-26372]
	_ = x[lbsMail-26373]
	_ = x[lbsManagerMessage-26374]
	_ = x[lbsLoginType-24848]
	_ = x[lbsConnectionID-24833]
	_ = x[lbsAskConnectionID-24834]
	_ = x[lbsWarningMessage-24835]
	_ = x[lbsRegulationHeader-26656]
	_ = x[lbsRegulationText-26657]
	_ = x[lbsRegulationFooter-26658]
	_ = x[lbsUserHandle-24849]
	_ = x[lbsUserRegist-24850]
	_ = x[lbsUserDecide-24851]
	_ = x[lbsAddProgress-24856]
	_ = x[lbsAskBattleResult-24864]
	_ = x[lbsAskGameVersion-24855]
	_ = x[lbsAskGameCode-24854]
	_ = x[lbsAskCountryCode-24853]
	_ = x[lbsAskPlatformCode-24852]
	_ = x[lbsAskKDDICharges-24898]
	_ = x[lbsPostGameParameter-24899]
	_ = x[lbsWinLose-24901]
	_ = x[lbsRankRanking-24900]
	_ = x[lbsDeviceData-24904]
	_ = x[lbsServerMoney-24905]
	_ = x[lbsAskNewsTag-26625]
	_ = x[lbsNewsText-26626]
	_ = x[lbsInvitationTag-26640]
	_ = x[lbsTopRankingTag-26705]
	_ = x[lbsTopRankingSuu-26706]
	_ = x[lbsTopRanking-26707]
	_ = x[lbsAskPatchData-26721]
	_ = x[lbsPatchHeader-26722]
	_ = x[lbsPatchData6863-26723]
	_ = x[lbsCalcDownloadChecksum-26724]
	_ = x[lbsPatchPing-26725]
	_ = x[lbsStartLobby-24897]
	_ = x[lbsPlazaMax-25091]
	_ = x[lbsPlazaTitle-25092]
	_ = x[lbsPlazaJoin-25093]
	_ = x[lbsPlazaStatus-25094]
	_ = x[lbsPlazaExplain-25098]
	_ = x[lbsPlazaEntry-25095]
	_ = x[lbsPlazaExit-25350]
	_ = x[lbsLobbyJoin-25347]
	_ = x[lbsLobbyEntry-25349]
	_ = x[lbsLobbyMatchingJoin-25615]
	_ = x[lbsLobbyExit-25608]
	_ = x[lbsRoomMax-25601]
	_ = x[lbsRoomTitle-25602]
	_ = x[lbsRoomStatus-25604]
	_ = x[lbsRoomCreate-25607]
	_ = x[lbsPutRoomName-26121]
	_ = x[lbsEndRoomCreate-26124]
	_ = x[lbsRoomEntry-25606]
	_ = x[lbsRoomExit-25857]
	_ = x[lbsRoomRemove-25861]
	_ = x[lbsPostChatMessage-26369]
	_ = x[lbsChatMessage-26370]
	_ = x[lbsUserSite-26371]
	_ = x[lbsLobbyRemove-25792]
	_ = x[lbsLobbyMatchingEntry-25614]
	_ = x[lbsWaitJoin-25862]
	_ = x[lbsMatchingEntry-25860]
	_ = x[lbsGoToTop-25096]
	_ = x[lbsReadyBattle-26896]
	_ = x[lbsAskMatchingJoin-26897]
	_ = x[lbsAskPlayerSide-26898]
	_ = x[lbsAskPlayerInfo-26899]
	_ = x[lbsAskRuleData-26900]
	_ = x[lbsAskBattleCode-26901]
	_ = x[lbsAskMcsAddress-26902]
	_ = x[lbsAskMcsVersion-26903]
	_ = x[lbsMatchingCancel-24581]
}

const _CmdID_name = "lbsLineChecklbsLogoutlbsShutDownlbsVSUserLostlbsMatchingCancellbsConnectionIDlbsAskConnectionIDlbsWarningMessagelbsLoginTypelbsUserHandlelbsUserRegistlbsUserDecidelbsAskPlatformCodelbsAskCountryCodelbsAskGameCodelbsAskGameVersionlbsAddProgresslbsAskBattleResultlbsStartLobbylbsAskKDDIChargeslbsPostGameParameterlbsRankRankinglbsWinLoselbsDeviceDatalbsServerMoneylbsPlazaMaxlbsPlazaTitlelbsPlazaJoinlbsPlazaStatuslbsPlazaEntrylbsGoToToplbsPlazaExplainlbsLobbyJoinlbsLobbyEntrylbsPlazaExitlbsRoomMaxlbsRoomTitlelbsRoomStatuslbsRoomEntrylbsRoomCreatelbsLobbyExitlbsLobbyMatchingEntrylbsLobbyMatchingJoinlbsLobbyRemovelbsRoomExitlbsMatchingEntrylbsRoomRemovelbsWaitJoinlbsPutRoomNamelbsEndRoomCreatelbsPostChatMessagelbsChatMessagelbsUserSitelbsSendMaillbsMaillbsManagerMessagelbsAskNewsTaglbsNewsTextlbsInvitationTaglbsRegulationHeaderlbsRegulationTextlbsRegulationFooterlbsTopRankingTaglbsTopRankingSuulbsTopRankinglbsAskPatchDatalbsPatchHeaderlbsPatchData6863lbsCalcDownloadChecksumlbsPatchPinglbsReadyBattlelbsAskMatchingJoinlbsAskPlayerSidelbsAskPlayerInfolbsAskRuleDatalbsAskBattleCodelbsAskMcsAddresslbsAskMcsVersion"

var _CmdID_map = map[CmdID]string{
	24577: _CmdID_name[0:12],
	24578: _CmdID_name[12:21],
	24579: _CmdID_name[21:32],
	24580: _CmdID_name[32:45],
	24581: _CmdID_name[45:62],
	24833: _CmdID_name[62:77],
	24834: _CmdID_name[77:95],
	24835: _CmdID_name[95:112],
	24848: _CmdID_name[112:124],
	24849: _CmdID_name[124:137],
	24850: _CmdID_name[137:150],
	24851: _CmdID_name[150:163],
	24852: _CmdID_name[163:181],
	24853: _CmdID_name[181:198],
	24854: _CmdID_name[198:212],
	24855: _CmdID_name[212:229],
	24856: _CmdID_name[229:243],
	24864: _CmdID_name[243:261],
	24897: _CmdID_name[261:274],
	24898: _CmdID_name[274:291],
	24899: _CmdID_name[291:311],
	24900: _CmdID_name[311:325],
	24901: _CmdID_name[325:335],
	24904: _CmdID_name[335:348],
	24905: _CmdID_name[348:362],
	25091: _CmdID_name[362:373],
	25092: _CmdID_name[373:386],
	25093: _CmdID_name[386:398],
	25094: _CmdID_name[398:412],
	25095: _CmdID_name[412:425],
	25096: _CmdID_name[425:435],
	25098: _CmdID_name[435:450],
	25347: _CmdID_name[450:462],
	25349: _CmdID_name[462:475],
	25350: _CmdID_name[475:487],
	25601: _CmdID_name[487:497],
	25602: _CmdID_name[497:509],
	25604: _CmdID_name[509:522],
	25606: _CmdID_name[522:534],
	25607: _CmdID_name[534:547],
	25608: _CmdID_name[547:559],
	25614: _CmdID_name[559:580],
	25615: _CmdID_name[580:600],
	25792: _CmdID_name[600:614],
	25857: _CmdID_name[614:625],
	25860: _CmdID_name[625:641],
	25861: _CmdID_name[641:654],
	25862: _CmdID_name[654:665],
	26121: _CmdID_name[665:679],
	26124: _CmdID_name[679:695],
	26369: _CmdID_name[695:713],
	26370: _CmdID_name[713:727],
	26371: _CmdID_name[727:738],
	26372: _CmdID_name[738:749],
	26373: _CmdID_name[749:756],
	26374: _CmdID_name[756:773],
	26625: _CmdID_name[773:786],
	26626: _CmdID_name[786:797],
	26640: _CmdID_name[797:813],
	26656: _CmdID_name[813:832],
	26657: _CmdID_name[832:849],
	26658: _CmdID_name[849:868],
	26705: _CmdID_name[868:884],
	26706: _CmdID_name[884:900],
	26707: _CmdID_name[900:913],
	26721: _CmdID_name[913:928],
	26722: _CmdID_name[928:942],
	26723: _CmdID_name[942:958],
	26724: _CmdID_name[958:981],
	26725: _CmdID_name[981:993],
	26896: _CmdID_name[993:1007],
	26897: _CmdID_name[1007:1025],
	26898: _CmdID_name[1025:1041],
	26899: _CmdID_name[1041:1057],
	26900: _CmdID_name[1057:1071],
	26901: _CmdID_name[1071:1087],
	26902: _CmdID_name[1087:1103],
	26903: _CmdID_name[1103:1119],
}

func (i CmdID) String() string {
	if str, ok := _CmdID_map[i]; ok {
		return str
	}
	return "CmdID(" + strconv.FormatInt(int64(i), 10) + ")"
}