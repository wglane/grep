# Go Grep Implementations

## Overview
This repository contains a lightweight demonstration of Grep in Go, showcasing five different implementations: `minimal`, `bad`, `ok`, `good`, and `great`. Each implementation varies in complexity and performance, providing a unique perspective on how the Grep functionality can be achieved in Go.

## Implementations
- **Minimal:** A very basic implementation with limited functionality.
- **Bad:** An example of how not to implement Grep, useful for learning purposes.
- **Ok:** A decent implementation with some improvements over the basic version.
- **Good:** A well-rounded implementation with better readability.
- **Great:** The most advanced implementation, using concurrency.

## Usage
To run any of these implementations, use the Go command-line tool with appropriate flags. The available options for each implementation can be viewed using the flag library.

### Example
An example of running the program with the `great` implementation is shown below:

```bash
go run grep.go -n -i -g great "heavens" texts/frankenstein.txt texts/iliad.txt
```

```
texts/frankenstein.txt:981:with frightful loudness from various quarters of the heavens. I remained,
texts/frankenstein.txt:1216:heavens; they have discovered how the blood circulates, and the nature of
texts/frankenstein.txt:1364:more certainly shine in the heavens than that which I now affirm is
texts/frankenstein.txt:2190:progress. It advanced; the heavens were clouded, and I soon felt the rain
texts/frankenstein.txt:2200:Switzerland, appeared at once in various parts of the heavens. The
texts/frankenstein.txt:3103:the mountain. The sun is yet high in the heavens; before it descends
texts/frankenstein.txt:3161:“Soon a gentle light stole over the heavens and gave me a sensation of
texts/frankenstein.txt:3549:not rain, as I found it was called when the heavens poured forth its
texts/frankenstein.txt:4358:mounted high in the heavens, but the cottagers did not appear. I
texts/frankenstein.txt:4405:dispersed the clouds that had loitered in the heavens; the blast tore
texts/frankenstein.txt:5568:my other sufferings. I looked on the heavens, which were covered by clouds
texts/frankenstein.txt:6069:“I am not mad,” I cried energetically; “the sun and the heavens, who
texts/frankenstein.txt:6343:The sun sank lower in the heavens; we passed the river Drance and
texts/frankenstein.txt:6370:in the west. The moon had reached her summit in the heavens and was
texts/frankenstein.txt:6373:scene of the busy heavens, rendered still busier by the restless waves
texts/frankenstein.txt:6699:invoked to aid me. Often, when all was dry, the heavens cloudless, and
texts/iliad.txt:1557:heavens: Virgil, like the same power in his benevolence, counselling
texts/iliad.txt:3496:Who in the heaven of heavens hast fixed thy throne,
texts/iliad.txt:4944:The power descending, and the heavens on fire!
texts/iliad.txt:6517:Such as the heavens produce: and round the gold
texts/iliad.txt:8240:The heavens attentive trembled as he spoke:[189]
texts/iliad.txt:9363:Heavens! how my country’s woes distract my mind,
texts/iliad.txt:12493:The heavens re-echo, and the shores reply:
texts/iliad.txt:12578:Heavens! what a prodigy these eyes survey,
texts/iliad.txt:13956:The god, whose lightning sets the heavens on fire,
texts/iliad.txt:15914:Then touch’d with grief, the weeping heavens distill’d
texts/iliad.txt:18997:Gives the loud signal, and the heavens reply.
texts/iliad.txt:19979:This the bright empress of the heavens survey’d,
texts/iliad.txt:24691:“Some of the epithets which Homer applies to the heavens seem to imply
texts/iliad.txt:24924:And beams of early light the heavens o’erspread.”
```