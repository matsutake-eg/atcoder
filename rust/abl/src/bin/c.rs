#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::fastout;
use proconio::input;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

#[allow(unused)]
const MOD: usize = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:usize,
        m:usize,
        ab: [(Usize1, Usize1); m],
    }

    let mut uft = unionfind::UnionFind::new(n);
    for (a, b) in ab {
        uft.union(a, b);
    }

    let mut size = HashMap::new();
    for i in 0..n {
        *size.entry(uft.find(i)).or_insert(0) += 1;
    }

    println!("{}", size.len() - 1);
}
