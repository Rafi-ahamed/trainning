	#Network means a group or system of interconnected people or things.

=internet is a worldwide network.
=when a network connected in local area that is local area network(sort form of [LAN])
=when computer connected in wide area that is called wide area network(sort form of[WAN]) 
=level of "OSI(Open System Interconnection)" has 7 levels= L7.compose email, L6:translation in machine code,L5: making a session layer,L4:making a transport layer for segmentation 
   ,L3:network(set on source ip or destination ip[packet]),L2:data link layer(set on "sender OR receiver MAC(Media Access Control)" that is called FRAME  ) 
   L1.physical layer(gets input as frame & output bits then converted & transfer in media[cats5 or cats6(copper) ,optical fiber, radio waves])

=every levels has protocol. L7.APPLICATION protocol is "HTTP" L6.translation layer protocol is "SSL" L5.session layer protocol is "SOCKET" L4.
    transport layer protocol is "TCP OR UDP" L3.network layer protocol is "IP" L2.DATALINK layer protocol is "ITEE802.1,802.11" L1.PHYSICAL has no "protocol"

=Protocol means a set of rules.It has two parts.There are:sender & receiver.when we make an application then who will send a request that is called client or sender & who will receive the message that is called server.example: web browser,FTP, Skype etc.

=how many hops comes in any website thats track command:"tracert" .In  every hops have a syntax :Hop RTT1 RTT2 RTT3 Name[ip address] that is called "RTT(Round trip time)". RTT  sometimes also referred to as latency.every "RTT" has a specific speed.

=New model of "OSI":l5:application-http,ftp
                    l4:transport-tcp,udp
                    l3:network-ip,ipv6
                    l2:datalink-802.2
                    l1:physical-no protocol  