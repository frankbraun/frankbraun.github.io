# Why I changed my mind on t-addresses in Zcash

*by [Frank Braun](/), 2025-07-11*

In Zcash there are two kind of addresses: transparent addresses (or t-addresses) and shielded addresses (or z-addresses).

Transparent addresses work very similar to addresses in Bitcoin. Shielded addresses employ zero-knowledge proofs in order to hide the sender, recipient, and amount of a transaction on the blockchain. The only information that remains visible for shielded transaction is that *a* transaction took place and nothing more.

All shielded transactions take place in a *shielded pool*, the current third iteration is called the *Orchard pool*, after the first pool called *Sprout* and the second one called *Sapling*.
For the purpose of this discussion we only consider the Orchard pool.

That gives us four transaction types:

- t -> t: *transparent* transactions similar to transactions in Bitcoin.
- s -> s: *shielded* transactions, totally opaque within a shielded pool.
- t -> s: a *shielding* transaction, moving coins from the transparent ledger into a shielded pool.
- s -> t: an *unshielding* transaction, moving coins from the shielded pool to the transparent ledger.

A common critique of Zcash is that it's not "private by default" (that is, it doesn't just have shielded transaction), in contrast to Monero, which is "private by default": All transactions within Monero have the same *type* and employ the same privacy features.

And that usually implies two things:

1. That Monero is better than Zcash, due to this "private by default" feature.
2. That Zcash should get rid of t-addresses and also become "private by default".

I used to utter the same reasoning, but now believe it to be wrong on both counts.

## A different perspective

> "A change in perspective is worth 80 IQ points."\
> — Alan Kay

I think the argument is best understood by looking at the situation from a different perspective and not conflating theory and practice.

So let's zoom out and consider what it means *in practice* to use a privacy coin like Zcash or Monero.

Until we reach a state where *all* transactions are private, using a privacy coin usually means a three step process:

1. Moving funds *into* the privacy coin.
2. Transacting *within* the privacy coin.
3. Moving funds *out* of the privacy coin.

Or in Zcash terms:

1. t -> s
2. s -> s
3. s -> t

The change of perspective is that Zcash is an *integrated* system of two separate parts: The shielded pool and a transparent ledger.

And just as Apple, by making both the hardware and the software, can make the better UX and better security features, Zcash can deliver better UX and better privacy by tightly integrating these two parts, **if executed correctly**.

So in order to properly compare Zcash and Monero from this vantage point, one really has to compare two things separately:

1. The privacy of shielded transactions in Zcash with the privacy of Monero transactions.
2. The integration into the wider crypto ecosystem of Zcash (considering t-addresses as well) and Monero.

## Privacy comparison

Much has been written about it, but I think it's fair to say that cryptographers agree that, just comparing the privacy of shielded transactions in Zcash with Monero transactions, Zcash has the better privacy properties.

From my understanding, Zcash offers you the theoretically possible maximum of privacy for on chain transactions and Monero doesn't. There are ongoing discussions of how much information can be derived from Monero transactions, what usage patterns are better for your privacy and which aren't, and so on.

I don't know about you, but when it comes to privacy in transactions, I just want the best possible and not think about potential weaknesses, how to use it best, and so on.

"I have the simplest tastes. I am always satisfied with the best."

The much quoted potential Zcash weakness of the trusted setup doesn't apply to the Orchard shielded pool anymore, so that's a non issue these days.

The much quoted Zcash weakness of not being "private by default" has been almost entirely removed by the [Zashi wallet](https://electriccoin.co/zashi/), see below.

I'm aware that Monero is working on an upgrade to Full-Chain Membership Proofs (FCMPs), called FCMP++, which will improve the privacy properties of Monero, but I'm considering here what is *available today*.

## Integration into the wider ecosystem

I believe that the release and wider adoption of the Zashi mobile wallet is a *game changer* for Zcash. Zcash has been described by some as the Apollo program of crypto and with the Zashi wallet in the ecosystem it's finally ready for lift-off.

So what does Zashi do that makes it such a game changer?

It supports signing of shielding, unshielding, and shielded transactions, but not transparent transactions. That is, it can send and receive both transparent and shielded transactions, but every received transparent UTXO first needs to be shielded before it can be spend again, either transparently (as an unshielding transaction) or privately (as a shielded transaction).

That means that all funds in Zashi are forced through the shielded pool.

So one could argue that it also makes Zcash "private by default", just with a larger part of the ecosystem directly integrated into the app. 

Having a good wallet is very important, because privacy comes more from the wallet than from the network or from your counterparties. A good wallet protects you from information leakage. Information about you that is never leaked from your local wallet in the first place cannot be leaked later by your counterparties or the network/blockchain.

It also integrates with the Keystone Hardware Wallet, which for the first time lets you *store* and _spend_ Orchard-pool Zcash on a hardware wallet. Other hardware wallets, like Trezor and Ledger, only support t-addresses. I agree with Zooko and others from the Zcash community that *privacy at rest* is even more important than *privacy in flight*. In short, the argument is that if you just use privacy in flight it's quite easy to deanonymize you by correlating amounts and timing transaction going in and out of a shielded pool. That is, the longer funds are in a shielded pool, privately held at rest, the harder it is deanonymize them, because they are "mixed" with more and more other funds. 

And with Zashi plus they Keystone hardware wallet integration it's now possible to both hold smaller (just in Zashi) and larger (combined with Keystone) ZEC amounts in a shielded pool, thereby increasing your own and everybody else's privacy at rest in the process.

As you can see on the [ZecHub Dashboard](https://zechub.wiki/dashboard), the amount of ZEC in the shielded pools has more than doubled in the last year and almost 14% of the circulating ZEC supply is held in the Orchard shielded pool now, trend increasing.

## Increased pluggability  and optional visibility

A lot in crypto is about the *pluggability* of a coin into the wider ecosystem. That's why things like EVM compatiblity matter for L1s and L2s, for example. And for Zcash having t-addresses and transparent transactions, which are very similar to Bitcoin addresses and transactions, means it's much easier to plug Zcash into the wider ecosystem. DEX integrations usually start with t-address support, because they basically just have to port the corresponding Bitcoin code, and can then later progress to implementing fully shielded support, like it's the case with [LeoDex](https://leodex.io).

Another example of a recent DEX adding support for ZEC with transparent addresses is [NEAR Intents](https://app.near-intents.org).

I believe this higher pluggablity of Zcash compared to Monero is the reason why there are more DEXs supporting Zcash than Monero and there is currently no cross-chain liquidity-pool DEX that supports XMR. And DEXs are very important for both privacy reasons (no KYC & AML nonsense) and usability (again, no KYC & AML nonsense).

Furthermore, t-addresses can also increase *visibility* for the use-cases where it actually makes sense. If I'm using a DEX, as a user I want to be as private as possible. But I want visibility into what's happening inside the DEX, because how else am I supposed to verify that the DEX is operating as advertised? So *optional visibility* is actually a feature and not a bug.

And just for completeness, having t-addresses allows for use-cases where one actually wants full transparency of the funds *at rest*. For example, using t-addresses for the cold wallet of exchanges or custody providers, allows easy *auditability* and *proof of funds*. 

## Conclusion

Given recent developments, mostly around Zashi, I changed my mind regarding t-addresses in Zcash. With what's already here (plus future work below) we can have:

- privacy by default, independent of third parties
- deeper integration into the wider crypto ecosystem (for example, with DEXs)
- optional visibility and auditability

Given a deeper integration into the wider crypto ecosystem (see below) and the arguments made above for the importance of privacy at rest there is also an argument to be made why value should accrue to ZEC the token: There is simply no other privacy coin in the market with both deep integration, optional visibility, and best in class privacy.

## Future work

Most importantly, Zashi currently only supports one fixed t-address, which is bad for privacy if t-addresses are used. Fixing that issue is on the Zashi roadmap, I argue on the [Zcash forum](https://forum.zcashcommunity.com/t/why-does-zashi-only-allow-1-transparent-address/51538/7) how I would design the UX in order to balance usability and privacy issues.

More options for storing and spending shielded Zcash on more widely used hardware wallets, like Trezor and Ledger, would be great. Until then, Keystone will get all the business for that use case and it will keep some Zcash users from storing ZEC long-term in a shielded pool. Personally, while I'm waiting for my Keystone wallet to arrive, I'm still storing most of my ZEC on a Trezor in a t-address.

Multisig support for hardware wallets, to be able to store very large amounts and for custodial setups with more than one signer, would also be very beneficial.

For good privacy in flight, transport layer privacy is also very important. It seem Zashi is rolling out [integrated Tor support](https://x.com/nate_zec/status/1942868170999828604), which is a good step in the right direction. In my opinion low-latency transport layer privacy without dummy traffic, like in Tor, doesn't cut it anymore and higher latency mixnets with dummy traffic, like Nym, are necessary. Until Nym releases their SDK, it cannot be integrated directly into Zashi, but for now the hardcore users have NymVPN in mixnet mode at their disposal.

Deeper integration of DEXs into Zashi (AFAIK that's being worked on) on both the receiving and sending side, to adhere to [the law of cryptocurrency isomorphism](https://juraj.bednar.io/en/blog-en/2022/10/24/the-law-of-cryptocurrency-isomorphism/), is also very much needed and should be deeply integrated into Zashi, for both receiving and sending arbitrary cryptocurrencies. ZEC will only benefit from that price wise, because there are good reasons to be made why one should actually **hold** ZEC, if one cares about ones privacy.

It's actually a much better value accrual story than can be made for Ethereum and Solana.
