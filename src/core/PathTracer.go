package main

import (
	"math"
	"math/rand"
)

var (
    room0 = Vec{-30, -0.5, -30}
    room1 = Vec{30, 15, 30}
    room2 = Vec{-20, 14, -20}
    room3 = Vec{20, 20, 20}
    plank0 = Vec{1.5, 18.5, -25}
    plank1 = Vec{6.5, 20, 25}
    dX = Vec{0.01, 0, 0}
    dY = Vec{0, 0.01, 0}
    dZ = Vec{0, 0, 0.01}
    lightDirection = (&Vec{0.3, 0.6, 0.4}).normalize()
    colorSun = Vec{50, 80, 100}
    colorWall = Vec{500, 400, 100}
)

func min(a float64, b float64) float64 {
    if a < b { 
        return a 
    } else { 
        return b 
    }
}

func carveOut(a float64, b float64) float64 {
    if b < 0 { return -b }
    return min(a, b)
}

type Hit int 
const (
    Wall Hit = iota
    Sun
    Figure
    None
)

/// Rectangle CSG equation. Returns minimum signed distance from
/// space carved by lowerLeft vertex and opposite rectangle vertex upperRight.
/// Negative return value if point is inside, positive if outside.
func probeBox(position *Vec, lowerLeft *Vec, upperRight *Vec) float64 {
    fromLowerLeft := position.minus(lowerLeft);
    toUpperRight := upperRight.minus(position);

    return -min(
                min(min(fromLowerLeft.X, toUpperRight.X),
                    min(fromLowerLeft.Y, toUpperRight.Y)),
                min(fromLowerLeft.Z, toUpperRight.Z));
}



func queryDatabase(position *Vec, hit *Hit) float64 {    
    distance := math.Max(1e9, position.X) - 0.5;
    *hit = Figure;
    plankedPosition := Vec{ math.Mod(math.Abs(position.X), 8.0), position.Y, position.Z };
    roomDist := min(
        -min(probeBox(position, &room0, &room1),
            probeBox(position, &room2, &room3)),
        probeBox(&plankedPosition, &plank0, &plank1));
    if roomDist < distance {
        distance = roomDist;
        *hit = Wall;
    }
    sun := 19.9 - position.Y;
    if sun < distance {
        *hit = Sun;
        return sun;
    }
    return distance;
}

func rayMarching(origin *Vec, direction *Vec, hitPos *Vec, hitNorm *Vec) Hit {
	var hitType Hit = None
	noHitCount := 0
	d := 0.0 // distance from the closest object in the world.
	for totalD := 0.0; totalD < 100.0; totalD += d {

		hitPos = origin.plus(direction.times(totalD))
		d = queryDatabase(hitPos, &hitType)
		if d < 0.01 || noHitCount > 99 {
			noHitCount += 1
			temp := None
			normX := queryDatabase(hitPos.plus(&dX), &temp) - d;
			normY := queryDatabase(hitPos.plus(&dY), &temp) - d;
			normZ := queryDatabase(hitPos.plus(&dZ), &temp) - d;
			hitNorm = (&Vec{ normX, normY, normZ }).normalize();
			return hitType;
		}
	}
	return None
}

func trace(origin *Vec, direction *Vec) *Vec {
	hitPoint := Vec { 0, 0, 0 }
	normal := Vec { 0, 0, 0 }
	result := Vec { 0, 0, 0 }
	attenuation := 1.0;
    newDirection :=  Vec{direction.X, direction.Y, direction.Z };
    newOrigin :=  &Vec{origin.X, origin.Y, origin.Z };
	for bounceCount := 3; bounceCount > 0; bounceCount -= 1 {
		hitType := rayMarching(newOrigin, &newDirection, &hitPoint, &normal);
		if hitType == None { break }
		if hitType == Figure {
			newDirection.minusM(normal.times(normal.dot(&newDirection) * 2.0));
			newOrigin = hitPoint.plus(newDirection.times(0.1));
			attenuation *= 0.2;
		} else if hitType == Wall {
			incidence := normal.dot(lightDirection)
			p := 6.283185 * rand.Float64();
			c := rand.Float64();
			s := math.Sqrt(1.0 - c);
			g := 1.0;
			if normal.Z < 0 { g = -1.0 };
			u := -1/(float64(g) + normal.Z);
			v := normal.X*normal.Y*u;
			a := Vec { v, g + normal.Y*normal.Y*u, -normal.Y };
			a.timesM(s * math.Cos(p));
			b := Vec { 1 + g*normal.X*normal.X*u, g*v, -g*normal.X }
			b.timesM(s * math.Sin(p));
			newDirection = a;
			newDirection.plusM(&b);
			newDirection.plusM(normal.times(math.Sqrt(c)));
			newOrigin = hitPoint.plus(newDirection.times(0.1))
			attenuation *= 0.2;
			ptAbove := hitPoint.plus(normal.times(0.1));
			if incidence > 0 {
				tmp := rayMarching(ptAbove, lightDirection, &hitPoint, &normal)
				if tmp == Sun {
					result.plusM(colorWall.times(attenuation*incidence));
				}
			}
		} else if hitType == Sun {
			result.plusM(colorSun.times(attenuation));
			break;
		}
	}
	return &result;
}

func run(position *Vec, dirObersver *Vec, samplesCount int, w int, h int) {
	dirLeft := (Vec{}).normalize()
}

public void run(Vec position, Vec dirObserver, int samplesCount, int w, int h) {
    Vec dirLeft = (new Vec(dirObserver.z, 0, -dirObserver.x)).normalize();
    //dirLeft.timesM(1.0 / w);
    dirLeft.timesM(1.0 / h);
    // Cross-product to get the up vector

    Vec dirUp = new Vec(dirObserver.y * dirLeft.z - dirObserver.z * dirLeft.y,
                        dirObserver.z * dirLeft.x - dirObserver.x * dirLeft.z,
                        dirObserver.x * dirLeft.y - dirObserver.y * dirLeft.x);
    dirUp.normalizeM();
    dirUp.timesM(1.0 / h);
    byte[] pixels = new byte[3 * w * h];

    Parallel.For (1, h - 1, (y) => {
        for (int x = w; x > 0; --x) {
            Vec color = new Vec(0, 0, 0);
            for (int p = samplesCount; p > 0; --p) {
                var randomLeft = dirLeft.times(x - w / 2 + rnd.NextDouble());
                var randomUp = dirUp.times((y - h / 2 + rnd.NextDouble()));
                var randomizedDir = new Vec(dirObserver.x, dirObserver.y, dirObserver.z);
                randomizedDir.plusM(randomLeft);
                randomizedDir.plusM(randomUp);
                randomizedDir.normalizeM();
                //Hit hType = Hit.None;
                var incr = trace(position, randomizedDir);
                if (y < h/2) {
                    ;
                }
                color.plusM(incr);
            }

            // Reinhard tone mapping
            color.timesM(241.0 / samplesCount);
            color = new Vec((color.x + 14.0) / (color.x + 255.0),
                            (color.y + 14.0) / (color.y + 255.0),
                            (color.z + 14.0) / (color.z + 255.0));
            color.timesM(255.0);
            int index = 3 * (w * y - w + x - 1);
            pixels[index    ] = (byte)color.x;
            pixels[index + 1] = (byte)color.y;
            pixels[index + 2] = (byte)color.z;
        }
    });
    Output.createBMP(pixels, w, h, "card.bmp");
}