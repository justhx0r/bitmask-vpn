// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p := messageKeyToIndex[key]
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"bn":    &dictionary{index: bnIndex, data: bnData},
		"en_GB": &dictionary{index: en_GBIndex, data: en_GBData},
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
		"es":    &dictionary{index: esIndex, data: esData},
		"es_ES": &dictionary{index: es_ESIndex, data: es_ESData},
		"eu":    &dictionary{index: euIndex, data: euData},
		"fr":    &dictionary{index: frIndex, data: frData},
		"it":    &dictionary{index: itIndex, data: itData},
		"lt":    &dictionary{index: ltIndex, data: ltData},
		"nl":    &dictionary{index: nlIndex, data: nlData},
		"pt_BR": &dictionary{index: pt_BRIndex, data: pt_BRData},
		"ru":    &dictionary{index: ruIndex, data: ruData},
		"sk":    &dictionary{index: skIndex, data: skData},
		"sv":    &dictionary{index: svIndex, data: svData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%[1]s is an easy, fast, and secure VPN service from %[2]s. %[1]s does not require a user account, keep logs, or track you in any way.\n\t    \nThis service is paid for entirely by donations from users like you. Please donate at %[3]s.\n\t\t\nBy using this application, you agree to the Terms of Service available at %[4]s. This service is provide as-is, without any warranty, and is intended for people who work to make the world a better place.\n\n\n%[1]v version: %[5]s": 2,
	"%s blocking internet": 24,
	"%s off":               19,
	"%s on":                18,
	"About":                3,
	"About...":             14,
	"An error has ocurred initializing the VPN: %v": 25,
	"Can't connect to %s: %v":                       7,
	"Cancel":                                        20,
	"Checking status...":                            9,
	"Connecting to %s":                              21,
	"Could not find a polkit authentication agent. Please run one and try again.": 5,
	"Donate":                       1,
	"Donate...":                    13,
	"Error starting VPN":           8,
	"Help...":                      12,
	"Initialization error":         4,
	"Missing authentication agent": 6,
	"Quit":                         15,
	"Reconnect":                    23,
	"Route traffic through":        16,
	"Stopping %s":                  22,
	"The %s service is expensive to run. Because we don't want to store personal information about you, there are no accounts or billing for this service. But if you want the service to continue, donate at least $5 each month.\n\t\nDo you want to donate now?": 0,
	"Turn off":          11,
	"Turn on":           10,
	"Use %s %v gateway": 17,
}

var bnIndex = []uint32{ // 27 elements
	0x00000000, 0x000000fe, 0x00000111, 0x00000601,
	0x0000061a, 0x0000063d, 0x00000731, 0x0000077f,
	0x000007d2, 0x00000822, 0x00000869, 0x0000087d,
	0x00000891, 0x000008aa, 0x000008c0, 0x000008dc,
	0x000008e9, 0x00000929, 0x0000096e, 0x00000981,
	0x00000994, 0x000009a4, 0x000009e4, 0x00000a07,
	0x00000a23, 0x00000a6b, 0x00000ad0,
} // Size: 132 bytes

const bnData string = "" + // Size: 2768 bytes
	"\x02The %[1]s service is expensive to run. Because we don't want to stor" +
	"e personal information about you, there are no accounts or billing for t" +
	"his service. But if you want the service to continue, donate at least $5" +
	" each month.\x0a\x09\x0aDo you want to donate now?\x02অনুদান\x02%[1]s %[" +
	"2]s থেকে একটি সহজ, দ্রুত এবং নিরাপদ ভিপিএন পরিষেবা। %[1]s ব্যবহারকারীর অ" +
	"্যাকাউন্টের প্রয়োজন নেই, লগ রাখে না, অথবা যে কোনও উপায়ে আপনাকে ট্র্য" +
	"াক করে।\x0a\x09    \x0aআপনার মতো ব্যবহারকারীদের কাছ থেকে অনুদান নিয়ে " +
	"এই সেবাটি সম্পূর্ণভাবে প্রদান করা হয় । অনুগ্রহ করে %[3]s এ দান করুন ।" +
	"\x0a\x09\x09\x0aএই অ্যাপ্লিকেশন ব্যবহার করে, আপনি %[4]s-এ উপলব্ধ পরিষেবা" +
	"র শর্তাবলীর সাথে সম্মত হন । এই পরিষেবাটি যেমন কোন ওয়ারেন্টি ছাড়াই প্" +
	"রদান করা হয়, এবং যারা বিশ্বকে একটি ভাল জায়গা করতে কাজ করে তাদের উদ্দ" +
	"েশ্যে ।\x0a\x0a\x0a%[1]v সংস্করণ: %[5]s\x02সম্পর্কে\x02সূচনা ত্রুটি" +
	"\x02পোলকিট প্রমাণীকরণ এজেন্ট খুঁজে পাওয়া যায়নি. অনুগ্রহ করে একটি চালান" +
	" এবং আবার চেষ্টা করুন ।\x02প্রমাণীকরণ এজেন্ট অনুপস্থিত\x02\u00a0%[1]s: %" +
	"[2]v সাথে সংযোগ করতে পারবেন না\x02ভিপিএন চালু করতে ত্রুটি হয়েছে\x02অবস্থ" +
	"া পরিক্ষা করা হচ্ছে ...\x02চালু কর\x02বন্ধ কর\x02সাহায্য...\x02অনুদান." +
	"..\x02সম্পর্কে...\x02বন্ধ\x02ট্রাফিক এর মাধ্যমে যাবে\x02%[1]s %[2]v গেটও" +
	"য়ে ব্যবহার করুন\x02%[1]s চালু\x02%[1]s বন্ধ\x02বাতিল\x02%[1]s এর সাথে" +
	" সংযুক্ত হচ্ছে\x02%[1]s বন্ধ হচ্ছে\x02পুনঃসংযোগ\x02%[1]s ইন্টারনেট প্রতি" +
	"রোধ করছে\x02VPN আরম্ভ করার সময় একটি ত্রুটি ঘটেছে: %[1]v"

var en_GBIndex = []uint32{ // 27 elements
	0x00000000, 0x000000fe, 0x00000105, 0x000002d3,
	0x000002d9, 0x000002ee, 0x0000033a, 0x00000357,
	0x00000375, 0x00000388, 0x0000039b, 0x000003a3,
	0x000003ac, 0x000003b4, 0x000003be, 0x000003c7,
	0x000003cc, 0x000003e2, 0x000003fa, 0x00000403,
	0x0000040d, 0x00000414, 0x00000428, 0x00000437,
	0x00000441, 0x00000459, 0x0000048a,
} // Size: 132 bytes

const en_GBData string = "" + // Size: 1162 bytes
	"\x02The %[1]s service is expensive to run. Because we don't want to stor" +
	"e personal information about you, there are no accounts or billing for t" +
	"his service. But if you want the service to continue, donate at least $5" +
	" each month.\x0a\x09\x0aDo you want to donate now?\x02Donate\x02%[1]s is" +
	" an easy, fast, and secure VPN service from %[2]s. %[1]s does not requir" +
	"e a user account, keep logs, or track you in any way.\x0a\x09    \x0aThi" +
	"s service is paid for entirely by donations from users like you. Please " +
	"donate at %[3]s.\x0a\x09\x09\x0aBy using this application, you agree to " +
	"the Terms of Service available at %[4]s. This service is provide as-is, " +
	"without any warranty, and is intended for people who work to make the wo" +
	"rld a better place.\x0a\x0a\x0a%[1]v version: %[5]s\x02About\x02Initiali" +
	"sation error\x02Could not find a Polkit authentication agent. Please run" +
	" one and try again.\x02Missing authentication agent\x02Can't connect to " +
	"%[1]s: %[2]v\x02Error starting VPN\x02Checking status...\x02Turn on\x02T" +
	"urn off\x02Help...\x02Donate...\x02About...\x02Quit\x02Route traffic thr" +
	"ough\x02Use %[1]s %[2]v gateway\x02%[1]s on\x02%[1]s off\x02Cancel\x02Co" +
	"nnecting to %[1]s\x02Stopping %[1]s\x02Reconnect\x02%[1]s blocking Inter" +
	"net\x02An error has ocurred initializing the VPN: %[1]v"

var en_USIndex = []uint32{ // 27 elements
	0x00000000, 0x000000fe, 0x00000105, 0x000002d3,
	0x000002d9, 0x000002ee, 0x0000033a, 0x00000357,
	0x00000375, 0x00000388, 0x0000039b, 0x000003a3,
	0x000003ac, 0x000003b4, 0x000003be, 0x000003c7,
	0x000003cc, 0x000003e2, 0x000003fa, 0x00000403,
	0x0000040d, 0x00000414, 0x00000428, 0x00000437,
	0x00000441, 0x00000459, 0x0000048a,
} // Size: 132 bytes

const en_USData string = "" + // Size: 1162 bytes
	"\x02The %[1]s service is expensive to run. Because we don't want to stor" +
	"e personal information about you, there are no accounts or billing for t" +
	"his service. But if you want the service to continue, donate at least $5" +
	" each month.\x0a\x09\x0aDo you want to donate now?\x02Donate\x02%[1]s is" +
	" an easy, fast, and secure VPN service from %[2]s. %[1]s does not requir" +
	"e a user account, keep logs, or track you in any way.\x0a\x09    \x0aThi" +
	"s service is paid for entirely by donations from users like you. Please " +
	"donate at %[3]s.\x0a\x09\x09\x0aBy using this application, you agree to " +
	"the Terms of Service available at %[4]s. This service is provide as-is, " +
	"without any warranty, and is intended for people who work to make the wo" +
	"rld a better place.\x0a\x0a\x0a%[1]v version: %[5]s\x02About\x02Initiali" +
	"zation error\x02Could not find a polkit authentication agent. Please run" +
	" one and try again.\x02Missing authentication agent\x02Can't connect to " +
	"%[1]s: %[2]v\x02Error starting VPN\x02Checking status...\x02Turn on\x02T" +
	"urn off\x02Help...\x02Donate...\x02About...\x02Quit\x02Route traffic thr" +
	"ough\x02Use %[1]s %[2]v gateway\x02%[1]s on\x02%[1]s off\x02Cancel\x02Co" +
	"nnecting to %[1]s\x02Stopping %[1]s\x02Reconnect\x02%[1]s blocking inter" +
	"net\x02An error has ocurred initializing the VPN: %[1]v"

var esIndex = []uint32{ // 27 elements
	0x00000000, 0x0000012a, 0x0000012f, 0x0000032e,
	0x00000338, 0x00000351, 0x000003b8, 0x000003da,
	0x000003fd, 0x00000415, 0x0000042e, 0x00000437,
	0x0000043e, 0x00000447, 0x0000044f, 0x0000045c,
	0x00000463, 0x00000476, 0x00000497, 0x000004a7,
	0x000004b5, 0x000004be, 0x000004d1, 0x000004e8,
	0x000004fa, 0x00000514, 0x00000545,
} // Size: 132 bytes

const esData string = "" + // Size: 1349 bytes
	"\x02El servicio %[1]s es costoso de mantener. Debido a que no queremos a" +
	"lmacenar información personal sobre los usuarios, no hay cuentas ni sist" +
	"ema de facturación en este servicio. Pero si deseas que el servicio cont" +
	"inúe, puedes donar al menos 5 USD al mes.\x0a\x09\x0a¿Quieres realizar l" +
	"a donación ahora?\x02Dona\x02%[1]s es un servicio de VPN fácil, rápido, " +
	"y seguro de %[2]s. %[1]s no requiere una cuenta de usuario, no mantiene " +
	"registros, ni te rastrea de ninguna manera.\x0a\x09    \x0aEste servicio" +
	" se paga en su totalidad por donaciones de usuarios como tú. Por favor d" +
	"ona en %[3]s.\x0a\x09\x09\x0aAl usar esta aplicación, aceptas los Términ" +
	"os del Servicio disponibles en %[4]s. Este servicio se provee tal cual, " +
	"sin garantía alguna, y está destinado a gente que trabaja para hacer del" +
	" mundo un lugar mejor.\x0a\x0a\x0aVersión de %[1]v: %[5]s\x02Acerca de" +
	"\x02Error de inicialización\x02No se pudo encontrar ningún agente de aut" +
	"enticacion de polkit. Por favor lanza uno y prueba de nuevo.\x02Falta un" +
	" agente de autenticación\x02No puedo conectar con %[1]s: %[2]v\x02Error " +
	"arrancando la VPN\x02Comprobando el estado...\x02Encender\x02Apagar\x02A" +
	"yuda...\x02Dona...\x02Acerca de...\x02Cerrar\x02Salir a través de\x02Usa" +
	" la salida de %[1]s por %[2]v\x02%[1]s encendida\x02%[1]s apagada\x02Can" +
	"celar\x02Conectando a %[1]s\x02Desconectando de %[1]s\x02Volver a conect" +
	"ar\x02%[1]s bloqueando internet\x02Ha ocurrido un error inicializando la" +
	" VPN: %[1]v"

var es_ESIndex = []uint32{ // 27 elements
	0x00000000, 0x000000fe, 0x00000104, 0x000002d2,
	0x000002dc, 0x000002f5, 0x00000356, 0x00000375,
	0x00000399, 0x000003ad, 0x000003c3, 0x000003cb,
	0x000003d6, 0x000003df, 0x000003e8, 0x000003f5,
	0x000003fb, 0x00000419, 0x00000439, 0x00000448,
	0x0000045a, 0x00000463, 0x00000476, 0x00000487,
	0x00000491, 0x000004ab, 0x000004dc,
} // Size: 132 bytes

const es_ESData string = "" + // Size: 1244 bytes
	"\x02The %[1]s service is expensive to run. Because we don't want to stor" +
	"e personal information about you, there are no accounts or billing for t" +
	"his service. But if you want the service to continue, donate at least $5" +
	" each month.\x0a\x09\x0aDo you want to donate now?\x02Donar\x02%[1]s is " +
	"an easy, fast, and secure VPN service from %[2]s. %[1]s does not require" +
	" a user account, keep logs, or track you in any way.\x0a\x09    \x0aThis" +
	" service is paid for entirely by donations from users like you. Please d" +
	"onate at %[3]s.\x0a\x09\x09\x0aBy using this application, you agree to t" +
	"he Terms of Service available at %[4]s. This service is provide as-is, w" +
	"ithout any warranty, and is intended for people who work to make the wor" +
	"ld a better place.\x0a\x0a\x0a%[1]v version: %[5]s\x02Acerca de\x02Error" +
	" de inicialización\x02No se pudo encontrar un agente de autenticación po" +
	"lkit. Por favor corre uno e intenta de nuevo.\x02Falta agente de autenti" +
	"cación\x02No se puede conectar a %[1]s: %[2]v\x02Error iniciando VPN\x02" +
	"Comprobando estado...\x02Activar\x02Desactivar\x02Ayuda...\x02Donar..." +
	"\x02Acerca de...\x02Salir\x02Enrutar tráfico a través de\x02Usar ruta de" +
	" salida %[1]s %[2]v\x02%[1]s activada\x02%[1]s desactivada\x02Cancelar" +
	"\x02Conectando a %[1]s\x02Deteniendo %[1]s\x02Reconnect\x02%[1]s bloquea" +
	"ndo Internet\x02An error has ocurred initializing the VPN: %[1]v"

var euIndex = []uint32{ // 27 elements
	0x00000000, 0x000000fe, 0x0000010d, 0x0000033c,
	0x00000347, 0x0000035c, 0x000003b7, 0x000003d7,
	0x000003fb, 0x00000415, 0x0000042b, 0x00000434,
	0x00000440, 0x0000044c, 0x0000045e, 0x0000046c,
	0x00000472, 0x00000494, 0x000004b1, 0x000004c2,
	0x000004d6, 0x000004db, 0x000004f2, 0x00000502,
	0x0000050f, 0x00000529, 0x00000558,
} // Size: 132 bytes

const euData string = "" + // Size: 1368 bytes
	"\x02The %[1]s service is expensive to run. Because we don't want to stor" +
	"e personal information about you, there are no accounts or billing for t" +
	"his service. But if you want the service to continue, donate at least $5" +
	" each month.\x0a\x09\x0aDo you want to donate now?\x02Dohaintza egin\x02" +
	"%[1]s erraza, azkarra eta segurua den VPN zerbitzua da, %[2]s(e)k eskain" +
	"ia. %[1]s(e)k ez du erabiltzaile kontu baten beharra, ez ditu datu priba" +
	"tuak gordetzen eta zaitu inolako jarraipenik egiten.\x0a\x09    \x0aZerb" +
	"itzu zu bezalako erabiltzaileen dohaintzez mantentzen da. Dohaintza egin" +
	" hemen, mesedez: %[3]s.\x0a\x09\x09\x0aAplikazio hau erabilita, %[4]s he" +
	"lbideko zerbitzu-baldintzak onartzen dituzu. Zerbitzu hau eskuratu duzun" +
	" bezala hornitzen da, inolako bermerik gabe, eta mundu hau leku hobe bat" +
	" bihurtzeko lan egiten dutenentzat bideratuta dago.\x0a\x0a\x0a%[1]v ber" +
	"tsioa: %[5]s\x02Honi buruz\x02Errorea habiaraztean\x02Ezin da polkit aut" +
	"entifikazio agente bat bilatu. Mesedez, exekutatu bat eta berria saiatu." +
	"\x02Autentifikazio agentea falta da\x02Ezin da %[1]s(e)ra konektatu: %[2" +
	"]v\x02Errorea VPNa habiaraztean\x02Egiaztatzen egoera...\x02Aktibatu\x02" +
	"Desaktibatu\x02Laguntza...\x02Dohaintza egin...\x02Honi buruz...\x02Irte" +
	"n\x02Bideratu trafikoa honen bitartez:\x02Erabili %[1]s %[2]v atebidea" +
	"\x02%[1]s aktibatuta\x02%[1]s desaktibatuta\x02Utzi\x02%[1]s(e)ra konekt" +
	"atzen\x02%[1]s gelditzen\x02Birkonektatu\x02%[1]s internet blokeatzen" +
	"\x02Errore bat gertatu da VPNa habiaraztean: %[1]v"

var frIndex = []uint32{ // 27 elements
	0x00000000, 0x00000154, 0x00000161, 0x000003a8,
	0x000003b2, 0x000003cc, 0x0000042e, 0x0000045a,
	0x00000487, 0x000004a3, 0x000004c1, 0x000004c9,
	0x000004d5, 0x000004dd, 0x000004ed, 0x000004fa,
	0x00000502, 0x0000051a, 0x0000053e, 0x00000550,
	0x00000566, 0x0000056e, 0x00000581, 0x00000591,
	0x000005a0, 0x000005b6, 0x000005fa,
} // Size: 132 bytes

const frData string = "" + // Size: 1530 bytes
	"\x02L’exploitation du service %[1]s coûte cher. Dans la mesure où ne nou" +
	"s voulons enregistrer aucun renseignement personnel à votre sujet, il n’" +
	"y ni compte ni facturation pour ce service. Mais si vous souhaitez toute" +
	"fois que le service continue, faites un don d’au moins 5\u00a0$ mensuell" +
	"ement\x0a\x09\x0aSouhaitez-vous faire un don maintenant\u2009?\x02Faire " +
	"un don\x02%[1]s est un service de RPV simple, rapide et sécurisé offert " +
	"par %[2]s. %[1]s ne demande pas de compte utilisateur, ne conserve pas d" +
	"e journaux, ni ne vous suit à la trace d’aucune façon.\x0a\x09    \x0aCe" +
	" service est entièrement financé par les dons d’utilisateurs comme vous." +
	" Veuillez faire un don sur %[3]s.\x0a\x09\x09\x0aEn utilisant cette appl" +
	"ication, vous acceptez les conditions générales d’utilisation qui se tro" +
	"uvent sur %[4]s. Ce service est fourni tel quel, sans aucune garantie, e" +
	"t s’adresse aux personnes qui œuvrent à rendre le monde meilleur.\x0a" +
	"\x0a\x0a%[1]v version \u00a0: %[5]s\x02À propos\x02Erreur d’initialisati" +
	"on\x02Impossible de trouver un agent d’authentification polkit. Veuillez" +
	" en exécuter un et ressayer.\x02L’agent d’authentification est manquant" +
	"\x02Impossible de se connecter à %[1]s\u00a0: %[2]v\x02Erreur du démarra" +
	"ge du RPV\x02Vérification de l’état…\x02Activer\x02Désactiver\x02Aide…" +
	"\x02Faire un don…\x02À propos…\x02Quitter\x02Acheminer le trafic par\x02" +
	"Utiliser la passerelle %[1]s %[2]v\x0a\x02%[1]s est activé\x02%[1]s est " +
	"désactivé\x02Annuler\x02Connexion à %[1]s\x02Arrêt de %[1]s\x02Se reconn" +
	"ecter\x02%[1]s bloque Internet\x02Une erreur est survenue lors de l’\x0a" +
	"initialisation du RPV\u00a0: %[1]v"

var itIndex = []uint32{ // 27 elements
	0x00000000, 0x000000fe, 0x00000103, 0x000002d1,
	0x000002de, 0x000002fb, 0x00000347, 0x00000369,
	0x00000387, 0x000003a3, 0x000003bc, 0x000003c4,
	0x000003cb, 0x000003d4, 0x000003dc, 0x000003ec,
	0x000003f1, 0x00000411, 0x00000429, 0x00000432,
	0x0000043c, 0x00000444, 0x00000458, 0x00000467,
	0x00000472, 0x0000048a, 0x000004bb,
} // Size: 132 bytes

const itData string = "" + // Size: 1211 bytes
	"\x02The %[1]s service is expensive to run. Because we don't want to stor" +
	"e personal information about you, there are no accounts or billing for t" +
	"his service. But if you want the service to continue, donate at least $5" +
	" each month.\x0a\x09\x0aDo you want to donate now?\x02Dona\x02%[1]s is a" +
	"n easy, fast, and secure VPN service from %[2]s. %[1]s does not require " +
	"a user account, keep logs, or track you in any way.\x0a\x09    \x0aThis " +
	"service is paid for entirely by donations from users like you. Please do" +
	"nate at %[3]s.\x0a\x09\x09\x0aBy using this application, you agree to th" +
	"e Terms of Service available at %[4]s. This service is provide as-is, wi" +
	"thout any warranty, and is intended for people who work to make the worl" +
	"d a better place.\x0a\x0a\x0a%[1]v version: %[5]s\x02Informazioni\x02Err" +
	"ore nell'inizializzazione\x02Could not find a polkit authentication agen" +
	"t. Please run one and try again.\x02Agente di autenticazione mancante" +
	"\x02Can't connect to %[1]s: %[2]v\x02Errore nell'avvio della VPN\x02Cont" +
	"rollando lo stato...\x02Accendi\x02Spegni\x02Aiuto...\x02Dona...\x02Info" +
	"rmazioni...\x02Esci\x02Instrada il traffico attraverso\x02Use %[1]s %[2]" +
	"v gateway\x02%[1]s on\x02%[1]s off\x02Annulla\x02Connettendo a %[1]s\x02" +
	"Fermando %[1]s\x02Riconnetti\x02%[1]s blocking internet\x02An error has " +
	"ocurred initializing the VPN: %[1]v"

var ltIndex = []uint32{ // 27 elements
	0x00000000, 0x0000013d, 0x00000146, 0x00000364,
	0x00000369, 0x0000037c, 0x000003e3, 0x00000408,
	0x00000431, 0x00000449, 0x0000045e, 0x00000467,
	0x00000471, 0x0000047c, 0x00000488, 0x00000490,
	0x00000498, 0x000004b5, 0x000004da, 0x000004e9,
	0x000004f9, 0x00000504, 0x0000051a, 0x00000529,
	0x0000053f, 0x00000559, 0x00000580,
} // Size: 132 bytes

const ltData string = "" + // Size: 1408 bytes
	"\x02%[1]s paslaugą brangiai kainuoja išlaikyti. Kadangi nenorime laikyti" +
	" jūsų asmeninės informacijos, todėl nėra jokių paskyrų ar apmokestinimo " +
	"už šią paslaugą. Tačiau, jei norite, kad paslauga būtų teikiama ir tolia" +
	"u, paaukokite kiekvieną mėnesį bent po $5 (JAV dolerius).\x0a\x09\x0aAr " +
	"norite paaukoti dabar?\x02Paaukoti\x02%[1]s yra lengva naudoti, greita i" +
	"r saugi VPN paslauga iš %[2]s. %[1]s nereikalauja naudotojo paskyros, ne" +
	"saugo žurnalų ir jokiu kitu būdu jūsų neseka.\x0a\x09    \x0aŠi paslauga" +
	" yra apmokama tik iš, tokių pačių naudotojų kaip jūs, paaukojimų. Prašom" +
	"e paaukokite, adresu %[3]s.\x0a\x09\x09\x0aNaudodami šią programą, sutin" +
	"kate su Naudojimosi Sąlygomis, kurias galite rasti adresu %[4]s. Ši pasl" +
	"auga yra teikiama esamu pavidalu, be jokių garantijų ir yra skirta žmonė" +
	"ms, kurie dirba, kad padarytų pasaulį geresnį.\x0a\x0a\x0a%[1]v versija:" +
	" %[5]s\x02Apie\x02Inicijavimo klaida\x02Nepavyko rasti politikos rinkini" +
	"o tapatybės nustatymo agento. Paleiskite jį ir bandykite dar kartą.\x02T" +
	"rūksta tapatybės nustatymo agento\x02Nepavyksta prisijungti prie %[1]s: " +
	"%[2]v\x02Klaida paleidžiant VPN\x02Tikrinama būsena...\x02Įjungti\x02Išj" +
	"ungti\x02Pagalba...\x02Paaukoti...\x02Apie...\x02Išeiti\x02Siųsti duomen" +
	"ų srautą per\x02Naudoti %[1]s %[2]v tinklų sietuvą\x02%[1]s įjungta\x02" +
	"%[1]s išjungta\x02Atsisakyti\x02Jungiamasi prie %[1]s\x02Stabdoma %[1]s" +
	"\x02Prisijungti iš naujo\x02%[1]s blokuoja internetą\x02Inicijuojant VPN" +
	", įvyko klaida: %[1]v"

var nlIndex = []uint32{ // 27 elements
	0x00000000, 0x000000fe, 0x00000106, 0x0000031e,
	0x00000323, 0x00000335, 0x00000384, 0x000003a1,
	0x000003c5, 0x000003e7, 0x000003fd, 0x00000409,
	0x00000416, 0x0000041e, 0x00000429, 0x00000431,
	0x00000439, 0x0000044c, 0x00000468, 0x00000472,
	0x0000047c, 0x00000486, 0x0000049a, 0x000004b0,
	0x000004ba, 0x000004d7, 0x00000521,
} // Size: 132 bytes

const nlData string = "" + // Size: 1313 bytes
	"\x02The %[1]s service is expensive to run. Because we don't want to stor" +
	"e personal information about you, there are no accounts or billing for t" +
	"his service. But if you want the service to continue, donate at least $5" +
	" each month.\x0a\x09\x0aDo you want to donate now?\x02Doneren\x02%[1]s i" +
	"s een gemakkelijke, snelle en veilige VPN-dienst van %[2]s. %[1]s vereis" +
	"t geen gebruikersaccount, houdt geen logboek bij en volgt u niet op wat " +
	"voor manier dan ook.\x0a\x09    \x0aDeze dienst wordt volledig gefinanci" +
	"erd door donaties van gebruikers zoals u. Gelieve bij te dragen op %[3]s" +
	".\x0a\x09\x09\x0aDoor deze applicatie te gebruiken, gaat u akkoord met o" +
	"nze gebruikersvoorwaarden die beschikbaar zijn op %[4]s. Deze dienst wor" +
	"dt geleverd zonder enige garantie en is bedoeld voor mensen die werken a" +
	"an een betere wereld. \x0a\x0a\x0a%[1]v version: %[5]s\x02Over\x02Initia" +
	"lisatiefout\x02Kan geen polkit authenticatieagent vinden. Voer er een ui" +
	"t en probeer opnieuw.\x02Authenticatieagent ontbreekt\x02Kan niet verbin" +
	"den met %[1]s: %[2]v\x02Fout bij het opstarten van de VPN\x02Status cont" +
	"roleren...\x02Inschakelen\x02Uitschakelen\x02Hulp...\x02Doneren...\x02Ov" +
	"er...\x02Stoppen\x02Stuur verkeer door\x02Gebruik %[1]s %[2]v gateway" +
	"\x02%[1]s aan\x02%[1]s uit\x02Annuleren\x02Verbinden met %[1]s\x02%[1]s " +
	"aan het stoppen\x02Reconnect\x02%[1]s blokkeert het internet\x02Er heeft" +
	" zich een fout voorgedaan bij het initialiseren van de VPN: %[1]v"

var pt_BRIndex = []uint32{ // 27 elements
	0x00000000, 0x000000fe, 0x00000111, 0x000002df,
	0x000002e5, 0x000002fd, 0x0000036e, 0x00000399,
	0x000003c3, 0x000003d9, 0x000003ef, 0x000003f5,
	0x000003fe, 0x00000407, 0x0000041d, 0x00000426,
	0x0000042b, 0x00000449, 0x00000467, 0x0000047a,
	0x00000490, 0x00000499, 0x000004ac, 0x000004ba,
	0x000004c4, 0x000004e6, 0x00000517,
} // Size: 132 bytes

const pt_BRData string = "" + // Size: 1303 bytes
	"\x02The %[1]s service is expensive to run. Because we don't want to stor" +
	"e personal information about you, there are no accounts or billing for t" +
	"his service. But if you want the service to continue, donate at least $5" +
	" each month.\x0a\x09\x0aDo you want to donate now?\x02Fazer uma doação" +
	"\x02%[1]s is an easy, fast, and secure VPN service from %[2]s. %[1]s doe" +
	"s not require a user account, keep logs, or track you in any way.\x0a" +
	"\x09    \x0aThis service is paid for entirely by donations from users li" +
	"ke you. Please donate at %[3]s.\x0a\x09\x09\x0aBy using this application" +
	", you agree to the Terms of Service available at %[4]s. This service is " +
	"provide as-is, without any warranty, and is intended for people who work" +
	" to make the world a better place.\x0a\x0a\x0a%[1]v version: %[5]s\x02So" +
	"bre\x02Erro na inicialização\x02Não foi possível encontrar um agente de " +
	"autenticação polkit. Por favor, execute um agente e tente novamente.\x02" +
	"Um agente de autenticação está faltando\x02Não é possível conectar a %[1" +
	"]s: %[2]v\x02Erro ao iniciar a VPN\x02Verificando estado...\x02Ligar\x02" +
	"Desligar\x02Ajuda...\x02Fazer uma doação...\x02Sobre...\x02Sair\x02Rotea" +
	"r o tráfego através de\x02Usar o gateway %[2]v de %[1]s\x02%[1]s está li" +
	"gada\x02%[1]s está desligada\x02Cancelar\x02Conectando a %[1]s\x02Parand" +
	"o %[1]s\x02Reconnect\x02%[1]s está bloqueando a Internet\x02An error has" +
	" ocurred initializing the VPN: %[1]v"

var ruIndex = []uint32{ // 27 elements
	0x00000000, 0x00000203, 0x0000022d, 0x000005d6,
	0x000005ee, 0x00000616, 0x000006c6, 0x00000705,
	0x00000742, 0x00000762, 0x00000785, 0x00000796,
	0x000007a9, 0x000007b9, 0x000007e6, 0x00000801,
	0x0000080c, 0x00000835, 0x00000863, 0x00000871,
	0x00000881, 0x00000892, 0x000008b2, 0x000008cb,
	0x000008e8, 0x00000916, 0x00000963,
} // Size: 132 bytes

const ruData string = "" + // Size: 2403 bytes
	"\x02Служба %[1]s платная. В связи с тем, что мы не хотим сохранять персо" +
	"нальную информацию о Вас, здесь нет регистрации или выставления счёта п" +
	"о этой службе. Но, если Вам хочется сохранить работу этой службы, жертв" +
	"уйте по крайней мере 5$ каждый месяц. Вы хотите пожертвовать в данный м" +
	"омент?\x02Сделать пожертвование\x02%[1]s это простой, быстрый и безопас" +
	"ный VPN-сервис от %[2]s. %[1]s не требует регистрации, не ведет журналы" +
	" или каким-либо другим способом не отслеживает вас.\x0a\x09    \x0aНаш с" +
	"ервис существует только на пожертвования от таких пользователей, как вы" +
	". Пожалуйста, поддержите нас тут %[3]s.\x0a\x09\x09\x0aИспользуя это при" +
	"ложение, вы соглашаетесь с Условиями использования, доступными по адрес" +
	"у %[4]s. Этот сервис предоставляется как есть, без каких-либо гарантий " +
	"и предназначен для людей, которые работают, чтобы сделать этот мир лучш" +
	"е.\x0a\x0a\x0a%[1]v версия: %[5]s\x02О приложении\x02Ошибка инициализац" +
	"ии\x02Не удалось найти агент проверки подлинности polkit. Пожалуйста, з" +
	"апустите один и попробуйте снова.\x02Отсутствует агент аутентификации" +
	"\x02Не удается подключиться к %[1]s: %[2]v\x02Ошибка запуска VPN\x02Пров" +
	"ерка статуса...\x02Включить\x02Выключить\x02Помощь...\x02Сделать пожерт" +
	"вование...\x02О приложении...\x02Выход\x02Маршрут трафика через\x02Испо" +
	"льзуется шлюз %[1]s %[2]v\x02%[1]s вкл.\x02%[1]s выкл.\x02Отменить\x02П" +
	"одключение к %[1]s\x02Остановка %[1]s\x02Переподключить\x02%[1]s блокир" +
	"овка интернета\x02Произошла ошибка при инициализации VPN: %[1]v"

var skIndex = []uint32{ // 27 elements
	0x00000000, 0x0000014a, 0x00000154, 0x00000374,
	0x0000037f, 0x00000398, 0x000003f4, 0x0000040a,
	0x00000440, 0x00000458, 0x0000046a, 0x00000473,
	0x0000047c, 0x00000489, 0x00000496, 0x000004a4,
	0x000004ae, 0x000004c3, 0x000004e1, 0x000004fe,
	0x0000051b, 0x00000524, 0x00000546, 0x00000562,
	0x00000572, 0x00000594, 0x000005bb,
} // Size: 132 bytes

const skData string = "" + // Size: 1467 bytes
	"\x02Služby aplikácie %[1]s sú nákladné na správu. Pretože o vás nechceme" +
	" zbierať osobné informácie, v tejto službe neexistujú žiadne účty ani po" +
	"platky. Ale ak chcete aby naša služba mohla fungovať aj v budúcnosti, pr" +
	"osíme vás, aby ste prispievali aspoň sumou 5 $ mesačne.\x0a\x0aChceli by" +
	" ste prispieť hneď teraz?\x02Prispieť\x02Aplikácia %[1]s predstavuje jed" +
	"noduché, rýchle a bezpečné VPN od poskytovateľa %[2]s. %[1]s nevyžaduje " +
	"používateľský účet a nijakým spôsobom vás nesleduje.\x0a\x0aTúto službu " +
	"financujú príspevky od používateľov, ako ste vy. Prosím, prispejte na %[" +
	"3]s.\x0a\x0aPoužívaním tejto aplikácie vyjadrujete súhlas s podmienkami " +
	"používania služby dostupnými na %[4]s. Táto služba je poskytovaná v stav" +
	"e ako je, bez akejkoľvek záruky a je určená pre ľudí, ktorý chcú zo svet" +
	"a urobiť lepšie miesto.\x0a\x0a\x0a%[1]v, verzia: %[5]s\x02O programe" +
	"\x02Chyba pri inicializácii\x02Nebolo možné nájsť overovacieho agenta po" +
	"lkit. Prosím, spustite ho a skúste to znova.\x02Chýba agent overenia\x02" +
	"Nebolo možné pripojiť sa k aplikácii %[1]s: %[2]v\x02Chyba pri spustení " +
	"VPN\x02Kontrola stavu...\x02Zapnúť\x02Vypnúť\x02Pomocník...\x02Prispieť." +
	"..\x02O programe...\x02Ukončiť\x02Smerovať prenos cez\x02Použiť prístup " +
	"%[1]s %[2]v\x02Aplikácia %[1]s je zapnutá\x02Aplikácia %[1]s je vypnutá" +
	"\x02Zrušiť\x02Pripájanie sa k aplikácii %[1]s\x02Zastavujem aplikáciu %[" +
	"1]s\x02Znova pripojiť\x02Aplikácia %[1]s blokuje internet\x02Pri spusten" +
	"í VPN nastala chyba: %[1]v"

var svIndex = []uint32{ // 27 elements
	0x00000000, 0x00000119, 0x00000120, 0x00000334,
	0x00000337, 0x0000034a, 0x000003a0, 0x000003bc,
	0x000003df, 0x000003f4, 0x0000040b, 0x00000412,
	0x0000041c, 0x00000426, 0x00000430, 0x00000436,
	0x0000043e, 0x00000453, 0x0000046f, 0x0000047d,
	0x0000048b, 0x00000492, 0x000004a6, 0x000004b4,
	0x000004c0, 0x000004d9, 0x00000506,
} // Size: 132 bytes

const svData string = "" + // Size: 1286 bytes
	"\x02%[1]s-tjänsten är dyr at driva. Då vi inte vill spara någon av din p" +
	"ersonliga information, så finns det inga konton eller fakturatjänster fö" +
	"r den här tjänsten. Men om du vill att tjänsten ska fortsätta fungera så" +
	" donera minst 5 dollar varje månad\x0a\x09\x0aVill du donera nu?\x02Done" +
	"ra\x02%[1]s är en enkel, snabb och säker VPN-tjänst från %[2]s. %[1]s be" +
	"höver inget användarkoto, behåller inga loggar och spårar dig inte på nå" +
	"got sätt.\x0a\x09    \x0aTjänsten betalas helt av donationer från använd" +
	"are som du. Vänligen donera på %[3]s.\x0a\x09\x09\x0aGenom att använda d" +
	"en här applikationen så godkänner du användarvillkoren, som finns tillgä" +
	"ngliga på %[4]s. Den här tjänsten erbjuds som den är, helt utan garantie" +
	"r och är till för människor som vill göra världen till en bättre plats." +
	"\x0a\x0a\x0a%[1]v version: %[5]s\x02Om\x02Fel vid initiering\x02Kunde in" +
	"te hitta polkit authentiseringsagent. Vänligen starta en och försök igen" +
	". \x02Saknar authentiseringsagent\x02Kan inte ansluta till %[1]s: %[2]v" +
	"\x02Fel vid start av VPN\x02Kontrollerar status...\x02Starta\x02Stäng av" +
	"\x02Hjälp...\x02Donera...\x02Om...\x02Avsluta\x02Skicka trafik igenom" +
	"\x02Använd %[1]s %[2]v gateway\x02%[1]s startad\x02%[1]s stoppad\x02Avbr" +
	"yt\x02Ansluter till %[1]s\x02Stoppar %[1]s\x02Återanslut\x02%[1]s blocke" +
	"rar internet\x02Ett fel uppstod vid initiering av VPN: %[1]v"

	// Total table size 22822 bytes (22KiB); checksum: D1112A41