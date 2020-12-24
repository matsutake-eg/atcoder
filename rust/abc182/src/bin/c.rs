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
        n:Chars,
    }

    let n_len = n.len();
    let mut cou = vec![0; 3];
    for c in n {
        let v = c as usize - '0' as usize;
        cou[v % 3] += 1;
    }

    let mut x = 0;
    for i in 0..3 {
        x += i * cou[i];
    }

    let mut ans = std::usize::MAX;
    for i in 0..=cou[1] {
        for j in 0..=cou[2] {
            if i + j == n_len {
                continue;
            }
            if (x - i * 1 - j * 2) % 3 == 0 {
                ans = min(ans, i + j);
            }
        }
    }
    if ans == std::usize::MAX {
        println!("-1");
        return;
    }
    println!("{}", ans);
}
