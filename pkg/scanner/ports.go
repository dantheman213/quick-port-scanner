package scanner

var Ports = map[int]string{
    5:     "rje",
    7:     "echo",
    11:    "systat",
    17:    "qotd",
    18:    "msp",
    19:    "chargen",
    20:    "FTP",
    21:    "FTP",
    22:    "SSH",
    23:    "Telnet",
    24:    "priv-mail",
    25:    "SMTP",
    26:    "RSFTP",
    37:    "TIME",
    39:    "rlp",
    42:    "nameserver",
    43:    "whois",
    49:    "tacacs",
    50:    "re-mail-ck",
    53:    "DNS",
    63:    "whois++",
    67:    "bootps,DHCP Server",
    68:    "bootpc,DHCP Client",
    69:    "tftp",
    70:    "gopher",
    79:    "finger",
    80:    "HTTP",
    88:    "Kerberos",
    95:    "supdup",
    101:   "hostname",
    102:   "iso-tsap",
    107:   "rtelnet",
    109:   "POP2",
    110:   "POP3",
    111:   "sunrpc",
    113:   "auth",
    115:   "sftp",
    117:   "uucp-path",
    119:   "USENET,nntp",
    123:   "NTP",
    135:   "MS EPMAP,DHCP,DNS,WINS,DCOM",
    136:   "sonos",
    137:   "NetBIOS,sonos",
    138:   "NetBIOS,sonos",
    139:   "NetBIOS,sonos",
    143:   "IMAP",
    161:   "SNMP",
    162:   "SNMP",
    163:   "cmip-man",
    164:   "cmip-agent",
    174:   "mailq",
    177:   "xdmcp",
    178:   "nextstep",
    179:   "BGP",
    191:   "prospero",
    194:   "IRC",
    199:   "SNMP,smux",
    201:   "AppleTalk",
    202:   "AppleTalk",
    204:   "AppleTalk",
    205:   "AppleTalk",
    206:   "AppleTalk",
    209:   "qmtp",
    210:   "z39.50",
    213:   "ipx",
    220:   "imap3",
    245:   "link",
    347:   "fatserv",
    363:   "rsvp_tunnel",
    369:   "rpc2portmap",
    389:   "LDAP",
    443:   "HTTPS",
    444:   "snpp",
    464:   "Kerberos,kpasswd",
    445:   "SMB,microsoft-ds",
    465:   "SMTPS",
    530:   "RPC",
    543:   "Kerberos,klogin",
    544:   "Kerberos,kshell",
    548:   "AFP",
    514:   "rsh",
    515:   "LPD",
    554:   "RTSP",
    631:   "IPP,CUPS",
    749:   "Kerberos,kadmin",
    760:   "Kerberos,kreg",
    765:   "webster",
    767:   "phonebook",
    853:   "DNS (TLS)",
    873:   "rsync",
    944:   "NFS",
    989:   "FTPS",
    990:   "FTPS",
    992:   "Telnet (TLS)",
    993:   "IMAP (TLS)",
    995:   "POP3 (TLS)",
    1025:  "NFS",
    1027:  "ICQ,NFS",
    1234:  "HTTP",
    1433:  "MS SQL Server",
    1434:  "MS SQL Server",
    1521:  "Oracle-RDS",
    1723:  "PPTP",
    1799:  "netrisk",
    1830:  "Oracle DB",
    1900:  "SSDP,UPnP,DLNA",
    2007:  "Elastic Graphics",
    2049:  "NFS",
    2181:  "ZooKeeper",
    2600:  "zebra",
    2888:  "ZooKeeper",
    3306:  "MySQL",
    3389:  "RDP",
    3401:  "sonos",
    3445:  "sonos",
    3500:  "sonos",
    3888:  "ZooKeeper",
    4070:  "sonos",
    4321:  "rwhois",
    4444:  "sonos",
    5000:  "UPnP,AirPrint",
    5002:  "rfe",
    5222:  "XMPP",
    5353:  "Bonjour,sonos",
    5432:  "PostgreSQL",
    5439:  "Redshift",
    5555:  "UPnP",
    5672:  "RabbitMQ",
    5800:  "RFB",
    5900:  "RFB",
    5984:  "CouchDB",
    5985:  "WinRM (HTTP)",
    5986:  "WinRM (HTTPS)",
    6000:  "x11",
    6379:  "Redis",
    6443:  "Kubernetes,HTTPS",
    6668:  "IRC",
    6881:  "BitTorrent",
    6882:  "BitTorrent",
    6883:  "BitTorrent",
    6884:  "BitTorrent",
    6885:  "BitTorrent",
    6886:  "BitTorrent",
    6887:  "BitTorrent",
    6888:  "BitTorrent",
    6889:  "BitTorrent",
    6969:  "BitTorrent Tracker,sonos",
    7000:  "Cassandra,AirPrint,AFS",
    7001:  "Cassandra,AFS",
    7002:  "AFS",
    7003:  "AFS",
    7004:  "AFS",
    7005:  "AFS",
    7006:  "AFS",
    7007:  "AFS",
    7008:  "AFS",
    7009:  "AFS",
    7010:  "AirPrint",
    7011:  "AirPrint",
    7100:  "AirPrint",
    8000:  "HTTP",
    8005:  "HTTP",
    8008:  "HTTP",
    8080:  "HTTP,Tomcat,Kubernetes",
    8091:  "HTTP",
    8092:  "HTTP",
    8093:  "HTTP",
    8200:  "DLNA,GoToMyPC,HTTP",
    8222:  "VMWare VI (HTTP)",
    8300:  "Consul",
    8301:  "Consul",
    8302:  "Consul",
    8333:  "Bitcoin",
    8400:  "Consul",
    8443:  "HTTPS",
    8500:  "Consul",
    8600:  "Consul",
    8612:  "AirPrint",
    8888:  "HTTP",
    9000:  "SonarQube,DBGp,SqueezeCenter,UDPCast,Play!,Hadoop,PHP-FPM",
    9001:  "MS SharePoint,ciscoxremote,Tor,DBGp Proxy,HSQLDB",
    9042:  "Cassandra",
    9043:  "WebSphere App Server",
    9060:  "WebSphere App Server",
    9080:  "WebSphere App Server (HTTP)",
    9092:  "H2/DBMS,Apache Kafka",
    9150:  "Tor Browser",
    9200:  "Elasticsearch",
    9222:  "Chrome DevTools",
    9229:  "NodeJS Remote Debug",
    9293:  "Sony PlayStation RemotePlay",
    9300:  "Elasticsearch",
    9333:  "LiteCoin",
    9339:  "Clash of Clans,game",
    9443:  "HTTPS,HTTP",
    9876:  "sd",
    10050: "Zabbix agent",
    10051: "Zabbix trapper",
    10110: "NMEA",
    11211: "memcached",
    11214: "memcached (TLS)",
    26000: "quake,game",
    27017: "MongoDB",
    33434: "traceroute",
    60000: "Spotify",
}
