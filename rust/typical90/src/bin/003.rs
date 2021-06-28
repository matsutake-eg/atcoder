#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::{fastout, input, marker::*};
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;
use superslice::Ext as _;

// CAUTION: don't print in 'match arm' because of fastout's bug
#[fastout]
fn main() {
    input! {
        n:usize,
        ab: [(Usize1, Usize1);n-1],
    }

    let mut g = vec![vec![]; n];
    for (a, b) in ab {
        g[a].push((b, 1));
        g[b].push((a, 1));
    }

    const INF: usize = std::usize::MAX;
    let mut ds = vec![INF; n];
    let x = 0;
    ds[x] = 0;
    let mut pq = BinaryHeap::new();
    pq.push((Reverse(0), x));
    while let Some((Reverse(d), v)) = pq.pop() {
        if d > ds[v] {
            continue;
        }
        for &(nx, c) in &g[v] {
            if d + c < ds[nx] {
                ds[nx] = d + c;
                pq.push((Reverse(ds[nx]), nx));
            }
        }
    }

    let max_x = ds.iter().max().unwrap();
    let x = ds.iter().position(|x| x == max_x).unwrap();
    let mut ds = vec![INF; n];
    ds[x] = 0;
    let mut pq = BinaryHeap::new();
    pq.push((Reverse(0), x));
    while let Some((Reverse(d), v)) = pq.pop() {
        if d > ds[v] {
            continue;
        }
        for &(nx, c) in &g[v] {
            if d + c < ds[nx] {
                ds[nx] = d + c;
                pq.push((Reverse(ds[nx]), nx));
            }
        }
    }
    println!("{}", ds.into_iter().max().unwrap() + 1);
}
