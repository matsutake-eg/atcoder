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
        a:[i64;n],
        q:usize,
        b:[i64;q],
    }

    let mut a = a;
    a.sort();

    for v in b {
        let idx = a.binary_search(&v).unwrap_or_else(|x| x);
        let mut ans = std::i64::MAX;
        if idx < n {
            ans = min(ans, (v - a[idx]).abs());
        }
        if idx + 1 < n {
            ans = min(ans, (v - a[idx + 1]).abs());
        }
        if idx >= 1 {
            ans = min(ans, (v - a[idx - 1]).abs());
        }
        println!("{}", ans);
    }
}
