# bdsm

Solidity smart contract environment tailored for fast smart contract development with minimal dependencies.

# Contracts

## factory/ERC20Factory.sol

An on-chain ERC20 factory designed for gas efficiency. Deploys the ERC20 contracts on-chain using runtime supplied arguments. The ERC20 source code is stored as a constant storage variable part of the factory contract, allowing for a cheaper than normal ERC20 factory,

## factory/RuntimeFactory.sol

A general factory contract that allows for deploying arbitrary contracts with runtime supplied bytecode.

## interfaces/ERC20I.sol

ERC20 interface

## math/SafeMath.sol

SafeMath library from openzeppelin

## ownership/Administration.sol

Basic contract ownership capabilities

## payments/eth_erc20_channel.sol

A contract allowing for the creation of unidirectional ETH and ERC20 payment channels, with the ability for micropayments (withdrawing from a channel without closing it).

## utils/Bytes.sol

Bytes manipulation contract

## utils/Address.sol

Address data type utilities