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
const MOD: i64 = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:i64,
    }

    let all = powm(10, n);
    let exclude = powm(9, n) * 2 % MOD;
    let duplicate = powm(8, n);
    println!("{}", ((all - exclude + MOD) % MOD + duplicate) % MOD);
}

fn powm(x: i64, y: i64) -> i64 {
    if y == 0 {
        return 1;
    } else if y % 2 == 0 {
        let x = powm(x, y / 2);
        return x * x % MOD;
    }
    x * powm(x, y - 1) % MOD
}
