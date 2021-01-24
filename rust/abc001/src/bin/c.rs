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
        deg:usize,
        dis:usize,
    }

    let w = match dis {
        0..=14 => 0,
        15..=92 => 1,
        93..=200 => 2,
        201..=326 => 3,
        327..=476 => 4,
        477..=644 => 5,
        645..=830 => 6,
        831..=1028 => 7,
        1029..=1244 => 8,
        1245..=1466 => 9,
        1467..=1706 => 10,
        1707..=1958 => 11,
        _ => 12,
    };
    let mut dir = match deg {
        113..=337 => "NNE",
        338..=562 => "NE",
        563..=787 => "ENE",
        788..=1012 => "E",
        1013..=1237 => "ESE",
        1238..=1462 => "SE",
        1463..=1687 => "SSE",
        1688..=1912 => "S",
        1913..=2137 => "SSW",
        2138..=2362 => "SW",
        2363..=2587 => "WSW",
        2588..=2812 => "W",
        2813..=3037 => "WNW",
        3038..=3262 => "NW",
        3263..=3487 => "NNW",
        _ => "N",
    };
    if w == 0 {
        dir = "C";
    }
    println!("{} {}", dir, w);
}
