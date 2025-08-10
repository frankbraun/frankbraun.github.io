# Freedom technology invariants

*by [Frank Braun](/), 2025-08-10*

This post is not about one particular freedom technology, but about what stays constant if everything changes — freedom technology invariants.

I usually define freedom technology, or *freedom tech*, as technology that is
able to give us more personal freedom in our lifetime. And since technology is
evolving very fast and the pace of change is likely to increase evermore it's
easy to lose track of what's really going to matter in the long term.

> "I very frequently get the question, 'What's going to change in the next ten
> years?' And that is a very interesting question; it's a very common one. I
> almost never get the question, 'What's not going to change in the next ten
> years?' And I submit to you that that second question is actually the more
> important of the two, because you can build a business strategy around the
> things that are stable in time."\
> — Jeff Bezos

For Amazon the things that are not going to change in next ten years are that
customers are going to want low prices, fast delivery, and a vast selection.

Let's apply this mental model to freedom tech.

- Who are the customers?
- What do they want?
- And what "invariants" can we derive from their wants regarding freedom
  technology?

## Who wants freedom?

Everybody is shopping, but not everybody wants to be free.

So while we sell goods to everybody, we cannot sell freedom to everybody.
Everybody is buying something, but not everybody is buying freedom. Some people
prefer other things, like safety, and are rather content with being controlled.
If that's wise is a different discussion.

> "Those who would give up essential Liberty, to purchase a little temporary
> Safety, deserve neither Liberty nor Safety."\
> — Benjamin Franklin

So what does freedom mean in the context of freedom technology?

**I think that, for a free society, the ability to communicate and transact
privately is essential.**

If you don't agree, ask yourself these questions:

- If you cannot communicate in private, how can you really be free?
- If you cannot transact in private, how can you really be free?

Let's assume for the sake of this discussion that our hypothetical "freedom
customer" wants to communicate and transact in private.

Or, more generally, let's assume that *free people want privacy*.

> We want privacy, because there is no freedom without it.

## The freedom tech stack

In order to have the ability to communicate and transact privately we need to
build out the *freedom tech stack*.

![Freedom tech stack](/asciiart/freedom-tech-stack.svg)

Let's go from the bottom to the top.

## Secure devices

Software can only ever be as secure as the endpoint, the device it is running
on, usually a computer or a mobile phone.

If you can compromise the [endpoint
security](https://en.wikipedia.org/wiki/Endpoint_security) you can compromise
all software that is running on it.

Therefore it is paramount that the hardware and software of our devices is
secure and *open*.

I think it's vastly underrated how dangerous a mandate of using [attested
applications](https://en.wikipedia.org/wiki/Trusted_Computing) for things like
age verification would be. In order for an application to be attestable it has
to run on an attestable device and that usually means you are running a mobile
phone which comes from Apple or Google with *their* OS on it and the app comes
from *their* app store. Which means you have neither control over the OS nor
over the software that you are running on it anymore.

> Invariant #1: We want to control our devices — the hardware itself and the
> software that runs on it.

## Secure software

Even if we control our devices the software on it is only as secure the the
entire software stack up and including the applications we run.

If there are exploitable bugs in the operating system or application software we
use then other security guarantees are moot.

So everything that improves operating system and programming language
correctness, including formal verification and tooling for better collaboration
and package delivery, is indirectly freedom tech.

> Invariant #2: We want our software to be open-source and correct.

## Anonymous messaging

The ability to communicate in private is necessary for freedom, but the secrecy
of the communication itself is not enough.

>"We kill people based on metadata."
> — Michael Hayden, former CIA & NSA director

How to achieve that is an implementation question, but keep in mind the
[Anonymity Trilemma](https://eprint.iacr.org/2017/954.pdf):

> "Strong anonymity, low bandwidth overhead, low latency — choose two"

But the implementation should be done with secure software running on a secure
device, every layer of the freedom tech stack is built on the layers below it.

> Invariant #3: We want to be able to communicate privately and anonymously.

## Digital cash

Money is defined as a medium of exchange, a unit of account, and a store of
value which should be portable, divisible, durable, rare, and fungible.

> Invariant #4:  We want digital cash — anonymous money that cannot be debased,
> censored, or traced.

## Nyms

Provable pseudonyms are *necessary* to build reputation systems.

Expensive pseudonyms would help against bots.

Pseudonyms allow to have an open discourse without the [chilling
efffect](https://en.wikipedia.org/wiki/Chilling_effect).

> Invariant #5: We want to be able to have provable  nyms — not tied to our
> legal identity.

## Free marketplaces

Free marketplaces, often wrongly named [darknet
markets](https://en.wikipedia.org/wiki/Darknet_market), require the entire
freedom tech stack discussed so far to function reliably and securely.

You cannot build a pyramid on quicksand.

> Invariant #6: We want to be able to transact in unrestricted market places — a
> free society requires free trade.

## Physical spaces

As physical beings we operate in the physical world.

That means in order to be truly free, floating freely in cyberspace is not going
to be enough.

Whoever controls your physical environment can control you.

> Invariant #7: We want to be able to control our physical environment.

## Conclusion

We want privacy, because there is no freedom without it.

And in order to get privacy we are building out the *freedom tech stack*.

When evaluating freedom tech we can ask ourselves these questions:

1. Does it lead to more control over our own computing devices?
2. Does it improve software openness and correctness?
3. Does it allow us to communicate more privately and anonymously?
4. Does it allow us to transact more freely, money wise?
5. Does it allow us to have nyms?
6. Does it allow us to transact more freely, in an unrestricted marketplace?
7. Does it give us more control over our physical environment?

The concrete technologies being built to achieve more freedom in our lifetimes
change all the time, but these principles remain constant — they are *freedom
technology invariants*.
