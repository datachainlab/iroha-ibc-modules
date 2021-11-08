#!/bin/python

import os
import sys
import binascii
from Crypto.Hash import keccak
from iroha import Iroha, IrohaCrypto, IrohaGrpc
from iroha.primitive_pb2 import can_call_engine

if sys.version_info[0] < 3:
    raise Exception('Python 3 or a more recent version is required.')

IROHA_HOST_ADDR = os.getenv('IROHA_HOST_ADDR', '127.0.0.1')
IROHA_PORT = os.getenv('IROHA_PORT', '50051')
ADMIN_ACCOUNT_ID = os.getenv('ADMIN_ACCOUNT_ID', 'admin@test')
ADMIN_PRIVATE_KEY = os.getenv(
        'ADMIN_PRIVATE_KEY',
        'f101537e319568c765b2cc89698325604991dca57b9716b58016b253506cab70'
)
TEST_ACCOUNT_ID = os.getenv('TEST_ACCOUNT_ID', 'test@test')
TEST_PRIVATE_KEY = os.getenv(
        'TEST_PRIVATE_KEY',
        '7e00405ece477bb6dd9b03a78eee4e708afc2f5bcdce399573a5958942f4a390'
)

iroha = Iroha(ADMIN_ACCOUNT_ID)
net = IrohaGrpc('{}:{}'.format(IROHA_HOST_ADDR, IROHA_PORT))

def trace(func):
    def tracer(*args, **kwargs):
        name = func.__name__
        print('\tEntering "{}"'.format(name))
        result = func(*args, **kwargs)
        print('\tLeaving "{}"'.format(name))
        return result

    return tracer

@trace
def send_transaction_and_print_status(transaction):
    hex_hash = binascii.hexlify(IrohaCrypto.hash(transaction))
    print('Transaction hash = {}, creator = {}'.format(
        hex_hash, transaction.payload.reduced_payload.creator_account_id))
    net.send_tx(transaction)
    for status in net.tx_status_stream(transaction):
        print(status)

@trace
def get_engine_receipts(transaction):
    tx_hash = binascii.hexlify(IrohaCrypto.hash(transaction))
    query = iroha.query('GetEngineReceipts', tx_hash=tx_hash)
    IrohaCrypto.sign_query(query, ADMIN_PRIVATE_KEY)
    response = net.send_query(query)
    return response.engine_receipts_response.engine_receipts

@trace
def deploy_contract(filepath):
    with open(filepath) as f:
        bc = f.read().strip()
    cmd = iroha.command('CallEngine', caller=ADMIN_ACCOUNT_ID, input=bc)
    tx = IrohaCrypto.sign_transaction(
            iroha.transaction([ cmd ]),
            ADMIN_PRIVATE_KEY
    )
    send_transaction_and_print_status(tx)
    return get_engine_receipts(tx)

@trace
def set_counter_address(callee, counter_address):
    k = keccak.new(digest_bits=256)
    k.update(b'setCounterAddress(address)')
    signature = k.digest()[:4].hex()
    counter_address = counter_address.zfill(64)
    input = signature + counter_address
    cmd = iroha.command('CallEngine', caller=ADMIN_ACCOUNT_ID, callee=callee, input=input)
    tx = IrohaCrypto.sign_transaction(
            iroha.transaction([ cmd ]),
            ADMIN_PRIVATE_KEY
    )
    send_transaction_and_print_status(tx)
    return get_engine_receipts(tx)

@trace
def incl(callee):
    k = keccak.new(digest_bits=256)
    k.update(b'incl()')
    signature = k.digest()[:4].hex()
    input = signature
    cmd = iroha.command('CallEngine', caller=ADMIN_ACCOUNT_ID, callee=callee, input=input)
    tx = IrohaCrypto.sign_transaction(
            iroha.transaction([ cmd ]),
            ADMIN_PRIVATE_KEY
    )
    send_transaction_and_print_status(tx)
    return get_engine_receipts(tx)

@trace
def decl(callee):
    k = keccak.new(digest_bits=256)
    k.update(b'decl()')
    signature = k.digest()[:4].hex()
    input = signature
    cmd = iroha.command('CallEngine', caller=ADMIN_ACCOUNT_ID, callee=callee, input=input)
    tx = IrohaCrypto.sign_transaction(
            iroha.transaction([ cmd ]),
            ADMIN_PRIVATE_KEY
    )
    send_transaction_and_print_status(tx)
    return get_engine_receipts(tx)

@trace
def call_library_internal(callee, x):
    k = keccak.new(digest_bits=256)
    k.update(b'callLibraryInternal(uint256)')
    signature = k.digest()[:4].hex()
    x = x.to_bytes(32, byteorder='big').hex()
    input = signature + x
    cmd = iroha.command('CallEngine', caller=ADMIN_ACCOUNT_ID, callee=callee, input=input)
    tx = IrohaCrypto.sign_transaction(
            iroha.transaction([ cmd ]),
            ADMIN_PRIVATE_KEY
    )
    send_transaction_and_print_status(tx)
    return get_engine_receipts(tx)

@trace
def call_library_public(callee, x):
    k = keccak.new(digest_bits=256)
    k.update(b'callLibraryPublic(uint256)')
    signature = k.digest()[:4].hex()
    x = x.to_bytes(32, byteorder='big').hex()
    input = signature + x
    cmd = iroha.command('CallEngine', caller=ADMIN_ACCOUNT_ID, callee=callee, input=input)
    tx = IrohaCrypto.sign_transaction(
            iroha.transaction([ cmd ]),
            ADMIN_PRIVATE_KEY
    )
    send_transaction_and_print_status(tx)
    return get_engine_receipts(tx)

@trace
def iroha_experiment_main(callee):
    k = keccak.new(digest_bits=256)
    k.update(b'main()')
    signature = k.digest()[:4].hex()
    cmd = iroha.command(
            'CallEngine',
            caller=ADMIN_ACCOUNT_ID,
            callee=callee,
            input=signature)
    tx = IrohaCrypto.sign_transaction(
            iroha.transaction([ cmd ]),
            ADMIN_PRIVATE_KEY)
    send_transaction_and_print_status(tx)
    return get_engine_receipts(tx)

@trace
def mint(callee, receiver_id, amount):
    k = keccak.new(digest_bits=256)
    k.update(b'mint(address,uint256)')
    signature = k.digest()[:4].hex()
    k = keccak.new(digest_bits=256)
    k.update(receiver_id.encode('ascii'))
    receiver = k.digest()[12:32].hex().zfill(64)
    amount = amount.to_bytes(32, byteorder='big').hex()
    input = signature + receiver + amount
    cmd = iroha.command('CallEngine', caller=ADMIN_ACCOUNT_ID, callee=callee, input=input)
    tx = IrohaCrypto.sign_transaction(
            iroha.transaction([ cmd ]),
            ADMIN_PRIVATE_KEY
    )
    send_transaction_and_print_status(tx)
    return get_engine_receipts(tx)

@trace
def transfer(callee, receiver_id, amount):
    k = keccak.new(digest_bits=256)
    k.update(b'send(address,uint256)')
    signature = k.digest()[:4].hex()
    k = keccak.new(digest_bits=256)
    k.update(receiver_id.encode('ascii'))
    receiver = k.digest()[12:32].hex().zfill(64)
    amount = amount.to_bytes(32, byteorder='big').hex()
    input = signature + receiver + amount
    cmd = iroha.command('CallEngine', caller=ADMIN_ACCOUNT_ID, callee=callee, input=input)
    tx = IrohaCrypto.sign_transaction(
            iroha.transaction([ cmd ]),
            ADMIN_PRIVATE_KEY
    )
    send_transaction_and_print_status(tx)
    return get_engine_receipts(tx)

#result = deploy_contract('./Coin.bytecode')
#callee = result[0].contract_address
#result = mint(callee, ADMIN_ACCOUNT_ID, 1000)
#print(result)
#result = transfer(callee, TEST_ACCOUNT_ID, 500)
#print(result)
#result = transfer(callee, TEST_ACCOUNT_ID, 300)
#print(result)
#result = transfer(callee, TEST_ACCOUNT_ID, 200)
#print(result)

if __name__ == '__main__':
    class Cli:
        def deploy_contract(self, bytecode_path):
            result = deploy_contract(bytecode_path)
            contract_address = result[0].contract_address
            print(contract_address)
        def set_counter_address(self, main_address, counter_address):
            result = set_counter_address(main_address, counter_address)
            call_result = result[0].call_result
            print(call_result.callee)
            print(call_result.result_data)
        def test(self, main_address):
            result = incl(main_address)
            log = result[0].logs[0]
            print(log.address)
            print(log.data)
            print(log.topics)
            result = incl(main_address)
            log = result[0].logs[0]
            print(log.address)
            print(log.data)
            print(log.topics)
            result = decl(main_address)
            log = result[0].logs[0]
            print(log.address)
            print(log.data)
            print(log.topics)
            result = call_library_internal(main_address, 3)
            log = result[0].logs[0]
            print(log.address)
            print(log.data)
            print(log.topics)
            result = call_library_public(main_address, 4)
            log = result[0].logs[0]
            print(log.address)
            print(log.data)
            print(log.topics)
        def iroha_experiment(self, iroha_experiment_address):
            results = iroha_experiment_main(iroha_experiment_address)
            for result in results:
                for log in result.logs:
                    print(log.address)
                    print(log.data)
                    print(log.topics)
    import fire
    fire.Fire(Cli)
