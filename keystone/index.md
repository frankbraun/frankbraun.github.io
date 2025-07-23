# Keystone hardware wallet test with shielded Zcash

*by [Frank Braun](https://frankbraun.org), 2025-07-23*

**Disclaimer: I don't have any relation with the Keystone company and paid for the hardware wallet myself.**

I finally got my hands on a Keystone 3 Pro hardware wallet in order to test it with the Zashi mobile wallet.

![Keystone hardware wallet](https://frankbraun.org/img/keystone.jpg)

After unboxing (nice experience, but we are not on YouTube here) I installed the Cypherpunk 2.0.4 [firmware](https://keyst.one/firmware) via SD card, in order to get shielded ZEC support. The same firmware also supports Monero (XMR) and Bitcoin (BTC) transactions, but I haven't tested those.

Creating a new 24 word seed phrase and writing it down was pretty straightforward. The mandatory verification comprises selecting the 24 words on the display in the right order which worked just fine, but I have seen better designs from a UX perspective. While the display and device has a very nice size and works well for normal operations, scrolling the word list up and down and selecting the words in order felt a bit clumsy to me.

After that hurdle has been taken, connecting the Keystone to Zashi was very straightforward. After that Zashi allows you to switch between the two "accounts" on the upper left (test was on iOS).

![Zashi account switcher](https://frankbraun.org/img/zashi.png)

Sending shielded ZEC to the Keystone is pretty straightforward. You can either
scan the QR code on the Keystone directly or just select the Keystone wallet
address in the list of recipients in Zashi (where it's automatically and
prominently put on the top).

## Signing a shielded transaction

Signing a shielded transaction on the Keystone is a simple process in Zashi:

1. Switch to the Keystone "account".
2. Prepare the transaction as with a normal transaction in Zashi.
3. Press `"Confirm with Keystone"`.
4. Scan the QR code with the Keystone.
5. Review and sign the transaction on the Keystone.
6. Press `"Get Signature"` on Zashi and scan the signed transaction QR code from the Keystone.

Zashi automatically transmits the transaction and *done*.

## Open issues

Shamir's Secret Sharing (SLIP39) seed phrases don't seem to work with ZEC on the Keystone, see [GitHub issue](https://github.com/KeystoneHQ/keystone3-firmware/issues/1806). There also seems to be some incompatibility with SLIP39 phrases imported from a Trezor, see [GitHub issue](https://github.com/KeystoneHQ/keystone3-firmware/issues/1506). I haven't tried these two options myself, but I consider proper and compatible SLIP39 support to be very important.

When entering the last number of a numerical PIN the gray dot doesn't turn orange before unlocking, see [GithHub issue](https://github.com/KeystoneHQ/keystone3-firmware/issues/1856). Not a big issue, but my obsessive brain doesn't like it.

"When you’re a carpenter making a beautiful chest of drawers, you’re not going to use a piece of plywood on the back, even though it faces the wall and nobody will ever see it. You’ll know it’s there, so you’re going to use a beautiful piece of wood on the back. For you to sleep well at night, the aesthetic, the quality, has to be carried all the way through."  
— Steve Jobs

## Conclusion

Having a **fully airgapped** [Keystone 3 Pro](https://keyst.one/shop/products/keystone-3-pro) hardware wallet that exclusively supports **shielded Zcash transactions** with a very nice integration into [Zashi](https://electriccoin.co/zashi/) **feels like magic**.

I'm looking forward to the upcoming Ledger support for shielded Zcash and I hope that Trezor revives their efforts to add shielded Zcash support. But having such a smooth experience transferring unsigned transactions to and signed transaction from the Keystone via QR codes has a totally different feel to it than fiddling with small buttons on a Ledger Nano X or Nano S Plus.

And the fully airgapped setup puts it on par with Bitcoin setups using hardware
wallets like Coinkite's Coldcard or Foundation's Passport.

I think that such hardware wallet support for *shielded Zcash transations* is absolutely necessary in order to set the flywheel in motion that I describe in my [Zcash investment thesis](https://frankbraun.org/zecbag/).

The cypherpunk future of unstoppable private money is here — it's just not very evenly distributed yet.
