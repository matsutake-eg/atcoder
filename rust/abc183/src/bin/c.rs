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
        k:usize,
        t:[[usize;n];n],
    }

    let mut ans = 0;
    for p in (1..n).permutations(n - 1).collect::<Vec<Vec<usize>>>() {
        let mut sum = t[0][p[0]] + t[p[p.len() - 1]][0];
        for i in 0..(p.len() - 1) {
            sum += t[p[i]][p[i + 1]];
        }
        if sum == k {
            ans += 1;
        }
    }
    println!("{}", ans);
}
