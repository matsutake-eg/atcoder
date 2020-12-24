#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:usize,
        t:Chars,
    }

    if n == 1 && t[0] == '1' {
        println!("{}", 20_000_000_000usize);
        return;
    }

    if [['1', '1', '0'], ['1', '0', '1'], ['0', '1', '1']]
        .iter()
        .any(|v| v.into_iter().cycle().zip(&t).all(|(a, b)| a == b))
    {
        let k = t.iter().filter(|&x| *x == '0').count();
        let mut ans = 10_000_000_000usize - k;
        if t[t.len() - 1] == '0' {
            ans += 1;
        }
        println!("{}", ans);
    } else {
        println!("{}", 0);
    }
}
