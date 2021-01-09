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
        h:[i64;n],
        w:[i64;m],
    }

    let mut h = h;
    h.sort();

    let mut d1 = vec![0; n];
    let mut i = 0;
    while i < d1.len() - 1 {
        d1[i] = (h[i] - h[i + 1]).abs();
        i += 2;
    }
    let d1 = d1.iter().cumsum().collect::<Vec<i64>>();

    let mut d2 = vec![0; n];
    let mut i = 1;
    while i < d2.len() - 1 {
        d2[i] = (h[i] - h[i + 1]).abs();
        i += 2;
    }
    let d2 = d2.iter().cumsum().collect::<Vec<i64>>();

    let mut ans = std::i64::MAX;
    for v in w {
        let mut d = 0;
        let idx = h.binary_search(&v).unwrap_or_else(|x| x);
        if idx % 2 == 0 {
            d += (v - h[idx]).abs();
            d += d2[n - 1] - d2[idx];
            if idx > 0 {
                d += d1[idx - 1];
            }
        } else {
            d += (v - h[idx - 1]).abs();
            d += d2[n - 1] - d2[idx - 1];
            if idx > 1 {
                d += d1[idx - 2];
            }
        }
        ans = min(ans, d);
    }
    println!("{}", ans);
}
