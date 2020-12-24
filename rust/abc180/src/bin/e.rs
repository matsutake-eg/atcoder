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
        xyz: [(i64,i64,i64);n],
    }

    let mut dist = vec![vec![0; n]; n];
    for i in 0..n {
        for j in 0..n {
            let mut now = (xyz[i].0 - xyz[j].0).abs();
            now += (xyz[i].1 - xyz[j].1).abs();
            now += max(0, xyz[j].2 - xyz[i].2);
            dist[i][j] = now as usize;
        }
    }

    let n2 = 1 << n;
    let mut dp = vec![vec![100_000_000_000; n]; n2];
    for i in 0..n {
        if i == 0 {
            continue;
        }
        dp[1 << i][i] = dist[0][i];
    }
    for i in 0..n2 {
        for j in 0..n {
            if i >> j & 1 != 1 {
                continue;
            }
            for k in 0..n {
                if i >> k & 1 == 1 {
                    continue;
                }
                dp[i | 1 << k][k] = min(dp[i | 1 << k][k], dp[i][j] + dist[j][k]);
            }
        }
    }
    println!("{}", dp[n2 - 1][0]);
}
