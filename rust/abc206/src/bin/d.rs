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
        a:[Usize1;n],
    }

    let a_max = *a.iter().max().unwrap();
    let mut uft = unionfind::UnionFind::new(a_max + 1);
    for i in 0..(n / 2) {
        if a[i] != a[n - 1 - i] {
            uft.union(a[i], a[n - 1 - i]);
        }
    }

    let mut size = HashMap::new();
    for i in 0..=a_max {
        *size.entry(uft.find(i)).or_insert(0) += 1;
    }
    let mut ans = 0;
    for (_, v) in size {
        ans += v - 1;
    }
    println!("{}", ans);
}
