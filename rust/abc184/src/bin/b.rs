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
        _:i64,
        x:i64,
        s:Chars,
    }

    let mut ans = x;
    for c in s {
        if c == 'o' {
            ans += 1;
        } else {
            ans = max(ans - 1, 0);
        }
    }
    println!("{}", ans);
}
