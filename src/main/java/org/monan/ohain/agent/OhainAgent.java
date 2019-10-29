package org.monan.ohain.agent;

import java.lang.instrument.ClassFileTransformer;
import java.lang.instrument.IllegalClassFormatException;
import java.lang.instrument.Instrumentation;
import java.security.ProtectionDomain;

/**
 * @author Yuewen.Huang
 * @email Yuewen.Huang@geely.com
 * @date 2019/10/29
 */
public class OhainAgent {

     public static void  premain(String agentOps, Instrumentation inst){
         inst.addTransformer(new ClassFileTransformer() {
             @Override
             public byte[] transform(ClassLoader loader, String className, Class<?> classBeingRedefined, ProtectionDomain protectionDomain, byte[] classfileBuffer) throws IllegalClassFormatException {

                 return new byte[0];
             }
         });

     }
}
