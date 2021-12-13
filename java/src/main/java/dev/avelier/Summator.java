package dev.avelier;

import org.graalvm.nativeimage.c.function.CEntryPoint;

public class Summator {

    private double sum;

    public double calc(double x) {
        sum += x;
        return sum;
    }

    @CEntryPoint(name = "Java_dev_avelier_callCalc")
    public static double callCalc(@CEntryPoint.IsolateThreadContext long isolateId, Summator s, double x) {
        return s.calc(x);
    }

    @CEntryPoint(name = "Java_dev_avelier_createSummator")
    public static Summator createSummator(@CEntryPoint.IsolateThreadContext long isolateId) {
        return new Summator();
    }

    @CEntryPoint(
        name = "Java_dev_avelier_createIsolate",
        builtin = CEntryPoint.Builtin.CREATE_ISOLATE)
    public static native long createIsolate();
}
