# Privacy needs to be baked into L1

*by [Frank Braun](/), 2025-07-28*

Privacy is the absence of information leakage and that implies it comes from value at rest, not value in motion. And privacy at rest can only be gained in a meaningful way by baking it into L1, not by adding it to higher layers.

That's a bold claim, so lets dig into it step by step.
## Privacy is the absence of information leakage

Privacy is not the presence of some mysterious quality, but the absence of information leakage. This is best understood if you look at the opposite of the term private — *public*.
If some information is public then anybody can potentially access that that information.
If some information is only known to a restricted group of people, then that information can still potentially be accessed by people outside of that group.

In the digital realm nothing can be taken back. If I send you a digital picture, I still have it.
I also can never control with certainty where this picture will end up afterwards.

"The only true secret is the one you take to the grave."  
— Anonymous

If you want to keep a secret, don't share it with anybody. If you share a secret with only one person and it becomes publicly known afterwards, you know who broke secrecy, but you cannot make it a secret again. If you share a secret with only two people and it becomes publicly known afterwards, you cannot even know anymore with certainty who broke secrecy, and you certainly cannot take it back. If you try to censor leaked information from the internet you will most likely become a victim of the [Streisand effect](https://en.wikipedia.org/wiki/Streisand_effect).

From a [threat model](https://en.wikipedia.org/wiki/Threat_model) perspective, it's best to assume that any information that is shared with anybody will become public knowledge.

Therefore, the only way to "gain" privacy is by not sharing information in the first place.

That is, privacy is the absence of information leakage.

## Privacy comes from value at rest, not value in motion

If you have a *store of value* (SoV) with perfect privacy — no information is leaked while value is stored or moved *within* the SoV — you only have privacy (no information leakage) while value is stored or moved within such a closed system.

The moment you move value *in* or *out* of the perfectly private SoV system, you leak some information and such information can typically be used to to derive more information.

And with the rapid improvements of AI technology it's best to assume that **any information that can be derived, will be derived**.

As Zooko explained in an [interesting thread on X](https://x.com/zooko/status/1935893105045389353), the best a perfectly private SoV can do is to **not** leak any information (my wording). Moving money in and out of it will **always** leak information.

![Privacy at rest infographic](/asciiart/privacy-at-rest.svg)

Sending funds *into* a perfectly private SoV will leak the time and value of the transaction to the public. Sending fund *out of* perfectly private SoV will also leak the time and value of the transaction to the public.

Using a perfectly private SoV as a [mixer](https://en.wikipedia.org/wiki/Cryptocurrency_tumbler) **won't work**, because transactions going in and coming out of it **can be correlated both in time and value**.

The only way to separate incoming and outgoing transactions is to separate them both in time and in value.

And the only way to do that is to use the perfectly private SoV actually as a store of value — a *savings account* — where the typical incoming and outgoing transaction is *significantly smaller* than the total amount of money saved. And an incoming transaction doesn't cause an outgoing transaction and vice versa (that is, there is no time correlation between incoming and outgoing transactions).

That's also the main argument in [my Zcash investment thesis](/zecbag) why value should accrue to privacy coins, if (and only if) the market demands more monetary privacy **and** understands this dynamic, but that's not the topic of this post.

Saving is not only necessary for capital accumulation, it's also a necessity for monetary privacy (unless you stay strictly within the perfectly private SoV).

Only the wealthy can afford to use non private payment systems with a certain assurance of privacy (as described above). If you don't have the necessary capital for a decently sized "private savings account" you **have to strictly transact within the perfectly private SoV** and never leave it in order to have monetary privacy.

That means, if you want to spread monetary privacy to everybody, the best way to do that is to onboard wealthy people with the use case of private SoV and not yet wealthy people with the use case of medium of exchange (MoE) in a **circular economy**.

## Privacy needs to be part of the base layer

Privacy is the absence of information leakage and that implies it comes from value at rest, not value in motion. And privacy at rest can only be gained in a meaningful way by baking it into L1, not by adding it to higher layers.

Now that we have discussed that privacy is the *absence of information leakage* and that *privacy comes from value at rest, not value in motion*, let's consider why it has to be baked into Layer 1 (L1), the base layer, of a cryptocurrency.

Before we get into details, here are a few well known examples of cryptocurrencies where privacy is baked into L1: Litecoin (LTC), Monero (XMR), and Zcash (ZEC).

And here are some well know crypto assets where privacy is **not** baked into L1: Bitcoin (BTC), Ethereum (ETH), and Solana (SOL).

I consider Decred (DCR) to be an edge case, which sits between these two categories, but for sake of brevity I'm not considering it in more detail here.

I'll explain in more detail below how I came to the distinction between these two categories (baked into L1 and privacy only in higher layers) and why I think that it's fundamental.

## Technology has no boundaries

In the category of the privacy coins that I mentioned where privacy is baked into L1, both Litecoin, since the MimbleWimble Extension Block (MWEB) extension, and Zcash (with both transparent and shielded transactions) cover all three parts (source, private SoV, and destination) of the graphic above. While Monero, with it's private by default design, covers only the middle part (private SoV). I explained in [another post](/t-addr) why I don't think that this matters for privacy, if you look at it from a different angle.

A great perspective on this (thanks to Nate Wilcox) is to zoom out to the scope of usage, not just "usage within this technological scope". For example, when ShapeShift was one of the most popular exchanges (before they were forced by regulators to introduce mandatory KYC), one of the most popular use of it was to exchange Bitcoin or Litecoin (before MWEB got introduced) for either Monero or Zcash and then exchange it back to the original coin, effectively using it as a mixer. This is a great demonstration that the existence of transparent addresses is irrelevant to this use case, and that no technology or service can constrain people's usage of it. It's also a great example that you cannot add privacy to value in motion, as discussed above.

The more successful these technologies are, the more people will use them to cross boundaries — technology, platform, product, social group, etc. That's a lot of the value of money! Boundary crossing. And if people are doing that, then details inside the technology, like optional transparent addresses vs. only shielded pools, are irrelevant. But "privacy comes from value at rest" remains relevant and becomes even more relevant.

Let's now look at the fundamental issues why privacy cannot be layered on top of L1 and then look at concrete, well known, examples where these issues can be witnessed in the wild.

## The problem of fungibility

The less deeply integrated privacy is into L1, the less fungible tokens become.

It's a total non-issue for Monero, because there is only one "pool". Zcash has deep integration and it's becoming normalized to use shielded pools, especially now with 19.5% of the circulating ZEC supply being held in shielded pools (increasing) and with wallets like [Zashi](https://electriccoin.co/zashi/) enforcing that all funds go through the shielded pool.

Litecoin is a bit of a special case, because after it launched in October 2011 it took over 10 years until MWEB was activated in May 2022. This also explains that only ~0.2%, namely [166,871 LTC](166,871)  out of a circulating supply of 76,117,814 LTC are in the MimbleWimble pool, which makes using MWEB on Litecoin more "suspicious" than using shielded pools in Zcash. But I expect this situation to improve, with more mobile wallet support being rolled out, [Cake Wallet](https://cakewallet.com/) was the first one to release MWEB support just in October 2024.

As a funny side story, Litecoin is only one of five crypto tokens that are *fully recognized* in DIFC (Dubai International Financial Centre), where the regulatory framework specifically prohibits privacy tokens. As it's often the case, crypto moves faster than even the most modern regulators.

But I digress, back on topic: In crypto assets like Bitcoin and Ethereum, if a user uses some form of mixer technology in an attempt to increase their privacy this is easily detected by now commonly used chain analysis technology, which makes the coins *less fungible*. If you don't believe me, try to pass a *source-of-funds* screening at a modern "crypto friendly" bank after you moved your entire stash through a mixer. Good luck. A lot of Bitcoin Maxis will have a rude awaking if they want to buy houses when BTC finally hits $1M and they realize that they listened too much to the "privacy experts" on their favorite podcast.

A crypto assets either has normalized privacy, like in Monero or Zcash, or the tokens are simply not fungible.

Of course, your banker could also complain about you using a privacy coin in the first place, that's a valid counterargument. However, we have to draw a line in the sand somewhere I still think it's both better for privacy and on principle to use systems which offer privacy *all the time*.

What would be more suspicious to you: Somebody using PGP to encrypt their emails (some of the time) or somebody using [Signal](https://signal.org)?

## The problem of custody

Besides the problem of fungibility, we also have the problem of custody. As we discussed above, *privacy comes from value at rest, not value in motion*. And in order to get privacy at rest you need to let your tokens rest. And if they rest in a (somewhat) custodial L2, then long rest periods become a problem, but long rest periods are exactly what you need.

While I cannot possibly give a comprehensive and in-depth review of all privacy technologies for Bitcoin and Ethereum, this is a good point in the story to give an overview of the most important ones and how they relate to the problem of custody.

The various CoinJoin technologies employed in Bitcoin, for example, in [Wasabi Wallet](https://wasabiwallet.io) and [Samourai Wallet](https://samouraiwallet.com) (website seized by U.S. law enforcement) are non-custodial, but also provide no privacy at rest. The famous Tornado Cash "mixer", employing zero-knowledge proofs in a smart contract on Ethereum, has some privacy at rest, but is non-custodial in nature (the employed smart contracts have no admin keys).

This following is going to get me a lot of hate, but here it comes. The Lightning Network (LN), the Manhattan Project of Bitcoin, was supposed to give us everything we ever wanted for Bitcoin: scalability and privacy, while still staying super decentralized and non-custodial. While being non-custodial (it just requires you to run your own lightning node with a hot wallet, what could possibly go wrong), it requires your wallet to be online in order to receive a payment. As of today there are 3,735.09 BTC locked in lightning channels (source: [1ML](https://1ml.com/statistics)), a whopping 0.017786% of the circulating supply.

The top 10 nodes hold ~ 64 % of the advertised channel capacity, with the single largest node (ACINQ) holding ~11 % (source also: [1ML](https://1ml.com/)), which is **exactly the hub and spoke network topology that early critics of LN predicted**. A criticism for which they were viciously attacked by Bitcoin Maxis. As much as one wishes a design to behave in a certain way, technical realities are technical realities and in the end reality always wins.

While it's great to pay for coffee (I really like [Phoenix wallet](https://phoenix.acinq.co) these days), I sincerely doubt that it will get us large scale privacy in Bitcoin, no matter how many bLIPs (Bitcoin Lightning Improvement Proposals) get activated. Sometimes it's just better to admit that one was climbing the wrong hill of technical complexity.

Let's now consider ecash systems, starting with the single mint design of [Cashu](https://cashu.space).

"**Cashu is not the second layer of Bitcoin** (or even the third), the pseudonymous author of the Cashu protocol Calle correctly says that the layer must support a sovereign transition to a lower layer in order to qualify for being another layer." (quoted by [Juraj Bednar](https://juraj.bednar.io/en/blog-en/2025/03/18/cashu-bitcoin-freebanking/))

I was one of the authors of the [Scrit whitepaper](https://github.com/scritcash/scrit-whitepaper/blob/master/scrit-whitepaper.pdf), a federated design that inspired [Fedimint](https://fedimint.org), and while the design spreads custody over multiple mints, it's still federated custody. 

"Fedimint is a federated custody protocol that complements the Bitcoin monetary protocol and Lightning payments protocol to provide a complete solution to holding, using, and securing Bitcoin at global scale."  
— [Obi Nwosu, Co-founder and CEO of Fedi Inc.](https://www.fedi.xyz/blog/fedi-raises-us-4-2-million-in-seed-financing)

While I still think that the various ecash designs, like Cashu and Fedi, hold great potential for fast, cheap, and private payments I now have sincere doubts about their ability to deliver *privacy at scale*, mainly due to the incompatibility between custody and value at rest: Their custodial nature (federated in the case of Fedi) requires the user to monitor them continuously and move funds to different mints, if necessary, which is in conflict with the requirement that privacy comes from value at rest.

## The problem of compliance

Which brings us to the final and largest issue with privacy on higher layers, the problem of compliance. Let's go through the same list of technologies and I think the picture will become pretty clear.

In April 2024 zkSNACKs Ltd, the company behind Wasabi Wallet, announced that U.S. citizens and residents were permanently blocked from accessing Wasabi Wallet and its services. And then in June 2024, the company officially shut down its CoinJoin coordinator service, probably to prevent a similar fate to that of Samourai Wallet.

In April 2024, U.S. authorities seized Samourai’s servers and domains and arrested co-founders Keonne Rodriguez and William Lonergan Hill, charging them with conspiracy to commit money laundering and operating an unlicensed money transmitting business (source: [Bitcoin Magazine](https://bitcoinmagazine.com/legal/freesamourai)). They are still awaiting trial, which is set to begin in November 2025.

And it continues with Tornado Cash which was put on the U.S. Treasury’s OFAC black list in August 2023, alleging it facilitated laundering of more than $7B. The protocol’s domain and GitHub repo were taken down, and multiple developers were arrested in the U.S. and Europe. One of the developers, Roman Storm, is currently dragged very publicly through U.S. courts (source: [Wikipedia](https://en.wikipedia.org/wiki/Tornado_Cash)).

There are more examples of mixers, but I think it should be clear by now that mixers are neither advisable for users, due to the lack of privacy at rest, nor for operators, due to "compliance challenges", no matter how decentralized the design appears — **even if there is no operator!**

While the specifics are murky, it's pretty clear that Lightning nodes could be considered *money transmitters* (MSB) in many cases, which would require corresponding licensing, especially if you hold customer funds or open channels for customers. To the best of my knowledge, this has never been formally brought to court, but the situation could change quickly if the LN would become a major method for "money laundering".

Cashu clearly is custodial and Calle is well served to stay pseudonymous. If a single mint would get significant volume it is very clear that the MSB definition would be applied **and enforced**, similar to more traditional mixers.

[Fedi](https://www.fedi.xyz), the company building products on top of Fedimint,
the protocol, is branded around *supporting communities*. While I know [elsirion](https://x.com/ericsirion), the main developer behind Fedimint, I have never spoken with him or other people from Fedi about this specific issue and I can only speculate that they are well aware of the compliance challenges revolving around operating a mint federation, which is why a branding and Go-to-Market strategy focused on communities running federations makes total sense to me. I certainly wouldn't want to run a company that runs federated ecash mints itself.

All these examples show that having privacy solutions at higher layers not only poses problems to the necessity for value at rest, but always comes with *significant compliance risks*, where **people regularly go to jail**.

It's all fun and games zapping each other on Nostr or paying toy amounts with ecash tokens, but for privacy at scale (certainly above the $1B threshold) **anything that even smells like custody or has any form of centralized control becomes a major compliance risk** of the "go to jail" category.

## Cognitive dissonance

A phrase often thrown around in Bitcoin circles is the goal of *separating money and state*. While I applaud that goal I don't think this goal can be achieved without *privacy at scale*: Are we really going to separate money and state with the state being able to watch every transaction?

The state has the monopoly of violence and the ability to print money out of thin air is an even better funding method for the state than taxing it's citizens, because debasement via money printing is a lot less obvious, and therefore less likely to face opposition, than direct taxation.

Bitcoin currently seems to be in the end stages of institutional adoption and I consider it very unlikely we will see any major privacy upgrades on L1, we can be lucky if we get the necessary quantum upgrades in time, given Bitcoin's ossification.

We need money that cannot be debased to escape inflation. But we need money that cannot be traced to escape confiscation. And as I argued in this post, this can only be achieved by baking privacy into L1.

I think this creates a lot of [cognitive dissonance](https://en.wikipedia.org/wiki/Cognitive_dissonance) in Bitcoiners. Many of us invested our money and sometimes our entire identity into the project.

Bitcoin is certainly the greatest *number go up (NgU) technology* ever invented, but at the same time the writing is on the wall: It will not give us monetary freedom and it will certainly not give us the separation of money and state. The process of institutionalization is likely to continue and in the end it will be just another asset, similar to gold, which will be held mostly by large institutions and whose **excellent monetary properties will help fund the exact thing a little longer that so many Bitcoiners despise — the state**.
## Conclusion

I argued that privacy is the absence of information leakage and that implies it comes from value at rest, not value in motion. And that privacy at rest can only be gained in a meaningful way by baking it into L1, not by adding it to higher layers, due to the problems of fungibility, custody, and compliance.

From the three cryptocurrencies mentioned in the beginning that have privacy baked into L1, Litecoin, Monero, and Zcash, only one currently has the necessary perfectly private SoV property — Zcash.

To me going down the Zcash rabbit hole definitely feels like Bitcoin early days again.

I gave [a talk](/bitcoin-in-the-counter-economy.pdf) at the first Bitcoin conference in London in 2012 about the dangers of Bitcoin being captured by states and banks. Now most of the technical issues in Bitcoin regarding privacy still persist and it's becoming completely institutionalized. Number will go up, sure, but it's not going to give us the freedoms we had hoped for, and it doesn't seem to give us [peer-to-peer electronic cash](https://bitcoin.org/bitcoin.pdf).

In order to succeed with the project of *separation of money and state* we need *privacy at scale* and that can only be achieved by baking privacy it into L1. And Zcash already did that with the necessary properties of being a perfectly private SoV — it's the best bet we have.

**There will be no third chance, it's now or never.**

So let's accept the reality of the situation, get over our cognitive dissonance and actually built on top of Zcash and [use it](/keystone) as a perfectly private SoV and MoE.

Bitcoin will continue to appreciate, but privacy at scale requires a more radical approach.

If you are the type that is convinced by big names, [Balaji Srinivasan](https://tim.blog/2021/03/25/balaji-srinivasan-transcript/), [Vitalik Buterin](https://x.com/VitalikButerin/status/993684262107926528) and [Tyler Winklevoss](https://coingape.com/tyler-winklevoss-extends-support-to-privacy-coin-project-zcash/) all recognize the importance of Zcash and it's unique position in the space.

Bitcoin has 3000x the market cap of Zcash, it's not going to falter if we go on a little side quest here — but maybe the main price really awaits us at the bottom of the shielded pool...

Thanks to Zooko Wilcox-O'Hearn and Juraj Bednar for input and feedback on this post.
