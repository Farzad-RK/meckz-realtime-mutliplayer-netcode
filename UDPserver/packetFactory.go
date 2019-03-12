package UDPserver

import (
	"github.com/google/flatbuffers/go"
	"meckz-netcode/UDPserver/ack"
	"meckz-netcode/UDPserver/conPackets"
	"meckz-netcode/UDPserver/packets"
)

func MakePacket(b *flatbuffers.Builder,sequenceNumber uint32,x float32 , y float32) []byte {
	// re-use the already-allocated Builder:
	b.Reset()

	position :=packets.CreateVec2(b,x,y)

	// write the Packet object:

	packets.PacketStart(b)
	packets.PacketAddPosition(b,position)
	packets.PacketAddSequenceNumber(b ,sequenceNumber)
	packets.PacketAddType(b,0)

	pkt := packets.PacketEnd(b)

	// finish the write operations by our User the root object:
	b.Finish(pkt)

	// return the byte slice containing encoded data:
	return b.Bytes[b.Head():]
}

func ReadPacket(buf []byte) (sequenceNumber uint32,position * packets.Vec2) {
	// initialize a User reader from the given buffer:
	pkt := packets.GetRootAsPacket(buf, 0)

	// point the name variable to the bytes containing the encoded name:

	position = pkt.Position(new(packets.Vec2))
	sequenceNumber = pkt.SequenceNumber()
	// copy the user's id (since this is just a uint64):

	return sequenceNumber ,position
}

func ReadConnPacket(buf [] byte)(token string ,sequenceNumber uint32){
	pkt := conPackets.GetRootAsConPacket(buf, 0)
	token = string(pkt.Token())
	sequenceNumber = pkt.SequenceNumber()
	return  token ,sequenceNumber
}
func MakeConnPacket(b *flatbuffers.Builder,sequenceNumber uint32,token string) []byte{
	// re-use the already-allocated Builder:
	b.Reset()
	_token := b.CreateByteString([]byte(token))
	// write the Packet object:
	conPackets.ConPacketStart(b)
	conPackets.ConPacketAddSequenceNumber(b,sequenceNumber)
	conPackets.ConPacketAddToken(b,_token)
	conPackets.ConPacketAddType(b,1)


	pkt := ack.AckEnd(b)
	// finish the write operations by our User the root object:
		b.Finish(pkt)

	// return the byte slice containing encoded data:
	return b.Bytes[b.Head():]
}
func MakeAck(b *flatbuffers.Builder,sequenceNumber uint32) []byte {
	// re-use the already-allocated Builder:
	b.Reset()

	ack.AckStart(b)
	ack.AckAddSequenceNumber(b,sequenceNumber)
	ack.AckAddType(b,2)
	pkt := ack.AckEnd(b)

	// finish the write operations by our User the root object:
	b.Finish(pkt)

	// return the byte slice containing encoded data:
	return b.Bytes[b.Head():]
}

func ReadAck(buf []byte)(sequenceNumber uint32)  {
	pkt :=ack.GetRootAsAck(buf,0)
	// point the name variable to the bytes containing the encoded name:
	sequenceNumber = pkt.SequenceNumber()
	// copy the user's id (since this is just a uint64):
	return sequenceNumber
}
func GetPacketType(buf [] byte) (packetType packets.Type)  {
	pkt := packets.GetRootAsPacket(buf,0)
	packetType = pkt.Type()
	return  packetType
}
