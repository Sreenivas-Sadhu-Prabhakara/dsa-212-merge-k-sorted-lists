# Merge K Sorted Lists

> NeetCode 150 · Linked List · Blind 75 — a min-heap of list heads, splicing the popped node straight into the output; nodes are reused, never copied. Go, stdlib only, tested.

## Problem

Merge k sorted linked lists into one sorted list — the node-splicing sibling of merge-k-sorted-arrays (dsa-073), famous enough as a hard-tier interview question to warrant its own treatment.

## Approach

The heap holds at most one node per list: each list's current head. Pop the smallest overall, splice it onto the output tail (`tail.Next = n`), and push that node's successor if any — the heap's `Less` compares node *values* while the heap stores *pointers*, so the spliced nodes carry their own identity and no values are copied. Nil lists are filtered before `heap.Init`. Terminating the tail explicitly (`tail.Next = nil`) matters: the last popped node may still point into its original list. The divide-and-conquer pairwise-merge alternative hits the same O(N log k) without a heap — worth naming, not worth both.

**Complexity:** O(N log k) time, O(k) heap.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
