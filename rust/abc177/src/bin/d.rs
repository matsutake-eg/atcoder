#![allow(unused_imports)]
use itertools::Itertools as _;
use itertools_num::ItertoolsNum as _;
use petgraph::unionfind::UnionFind;
use proconio::fastout;
use proconio::input;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;
use std::iter;

#[allow(unused)]
const INF: usize = std::usize::MAX / 4;
#[allow(unused)]
const MOD: usize = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[fastout]
fn main() {
    input! {
        n:usize,
        m:usize,
        ab:[(Usize1,Usize1);m],
    }

    let mut uft = UnionFind::new(n);
    for (a, b) in ab {
        uft.union(a, b);
    }

    let mut size = HashMap::new();
    for i in 0..n {
        *size.entry(uft.find(i)).or_insert(0) += 1;
    }
    println!("{}", size.values().max().unwrap());
}
