package org.monan.ohain.proxy;

import javassist.CannotCompileException;
import javassist.ClassPool;
import javassist.CtClass;
import javassist.CtMethod;

import java.lang.reflect.Method;
import java.lang.reflect.Parameter;

/**
 * @author Yuewen.Huang
 * @email Yuewen.Huang@geely.com
 * @date 2019/10/29
 */
public class Proxy {

    public Class getProxy(Class source){
        ClassPool cp = ClassPool.getDefault() ;

        String className = source.getName()+"$Proxy" ;
        CtClass ctClass = cp.makeClass(className) ;

        Method[] methods = source.getDeclaredMethods() ;
        for(Method method : methods) {
            String src = "" ;
            if ("void".equals(method.getReturnType().getSimpleName())){

            } else {

            }

            CtMethod cm = null;
            try {
                cm = CtMethod.make("",ctClass);
                ctClass.addMethod(cm);
            } catch (CannotCompileException e) {
                throw new RuntimeException(e) ;
            }

        }

        Class clazz = null ;
        try {
             clazz = ctClass.toClass() ;
        } catch (CannotCompileException e) {
            throw new RuntimeException(e) ;
        }
        return clazz ;
    }


    private String getExceptions(){
        return null ;
    }

    private String getParams(Parameter[] parameters){
        StringBuilder params = new StringBuilder() ;
        for (int index = 0 ,max = parameters.length ; index < max ;index ++) {
            params.append(parameters[index].getClass().getName()).append(" " ).append("$"+index).append(",");
        }

        params.deleteCharAt(params.length()-1) ;
        return params.toString() ;
    }

    private String getModifiers(int modifier){
        String modifierName = "public" ;
        if (modifier == 2 ){
            modifierName = "private" ;
        } else if(modifier == 4){
            modifierName = "protected" ;
        }
        return modifierName ;
    }
}
