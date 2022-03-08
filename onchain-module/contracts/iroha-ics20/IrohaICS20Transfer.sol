// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "@hyperledger-labs/yui-ibc-solidity/contracts/core/types/Channel.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCModule.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCHandler.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCHost.sol";
import "./IIrohaICS20Transfer.sol";
import "./IrohaAssetPacketData.sol";
import "./IrohaUtil.sol";
import "../old-experiments/IrohaApi.sol";

abstract contract IrohaICS20Transfer is IIrohaICS20Transfer {

    IBCHandler ibcHandler;
    IBCHost ibcHost;

    constructor(IBCHost host_, IBCHandler ibcHandler_) {
        ibcHost = host_;
        ibcHandler = ibcHandler_;
    }

    // INFO: This function was originally external, but changed to public, after encountering "Stack too deep".
    function sendTransfer(
        string memory srcAccountId,
        string memory destAccountId,
        string memory assetId,
        string memory description,
        string memory amount,
        string memory sourcePort,
        string memory sourceChannel,
        uint64 timeoutHeight
    ) public override {
        require(burn(srcAccountId, assetId, _makeDescription("burn", sourcePort, sourceChannel), amount), "burn failed");

        _sendPacket(
            IrohaAssetPacketData.Data({
                src_account_id: srcAccountId,
                dest_account_id: destAccountId,
                asset_id: assetId,
                description: description,
                amount: amount
            }),
            sourcePort,
            sourceChannel,
            timeoutHeight
        );
    }

    /// Module callbacks ///

    function onRecvPacket(Packet.Data calldata packet) external override returns (bytes memory acknowledgement) {
        IrohaAssetPacketData.Data memory data = IrohaAssetPacketData.decode(packet.data);
        return _newAcknowledgement(mint(data.dest_account_id, data.asset_id, data.description, data.amount));
    }

    function onAcknowledgementPacket(Packet.Data calldata packet, bytes calldata acknowledgement) external override {
        if (!_isSuccessAcknowledgement(acknowledgement)) {
            _refundTokens(IrohaAssetPacketData.decode(packet.data), packet.source_port, packet.source_channel);
        }
    }

    function onChanOpenInit(Channel.Order, string[] calldata, string calldata, string calldata channelId, ChannelCounterparty.Data calldata, string calldata) external override {}
    function onChanOpenTry(Channel.Order, string[] calldata, string calldata, string calldata channelId, ChannelCounterparty.Data calldata, string calldata, string calldata) external override {}
    function onChanOpenAck(string calldata portId, string calldata channelId, string calldata counterpartyVersion) external override {}
    function onChanOpenConfirm(string calldata portId, string calldata channelId) external override {}
    function onChanCloseInit(string calldata portId, string calldata channelId) external override {}
    function onChanCloseConfirm(string calldata portId, string calldata channelId) external override {}

    /// Virtual internal functions ///

    function burn(string memory srcAccountId, string memory assetId, string memory description, string memory amount) virtual internal returns (bool);
    function mint(string memory destAccountId, string memory assetId, string memory description, string memory amount) virtual internal returns (bool);

    /// Private functions ///

    function _sendPacket(IrohaAssetPacketData.Data memory data, string memory sourcePort, string memory sourceChannel, uint64 timeoutHeight) private {
        (Channel.Data memory channel, bool found) = ibcHost.getChannel(sourcePort, sourceChannel);
        require(found, "channel not found");
        ibcHandler.sendPacket(Packet.Data({
            sequence: ibcHost.getNextSequenceSend(sourcePort, sourceChannel),
            source_port: sourcePort,
            source_channel: sourceChannel,
            destination_port: channel.counterparty.port_id,
            destination_channel: channel.counterparty.channel_id,
            data: IrohaAssetPacketData.encode(data),
            timeout_height: Height.Data({revision_number: 0, revision_height: timeoutHeight}),
            timeout_timestamp: 0
        }));
    }

    function _newAcknowledgement(bool success) private pure returns (bytes memory) {
        bytes memory acknowledgement = new bytes(1);
        if (success) {
            acknowledgement[0] = 0x01;
        } else {
            acknowledgement[0] = 0x00;
        }
        return acknowledgement;
    }
    
    function _isSuccessAcknowledgement(bytes memory acknowledgement) private pure returns (bool) {
        require(acknowledgement.length == 1);
        return acknowledgement[0] == 0x01;
    }

    function _refundTokens(IrohaAssetPacketData.Data memory data, string memory sourcePort, string memory sourceChannel) private {
        require(mint(data.src_account_id, data.asset_id, _makeDescription("refund", sourcePort, sourceChannel), data.amount));
    }

    function _makeDescription(string memory operation, string memory sourcePort, string memory sourceChannel) private returns (string memory) {
        return string(abi.encodePacked(
            "operation=",
            operation,
            "sourcePort=",
            sourcePort,
            "sourceChannel",
            sourceChannel
        ));
    }

}