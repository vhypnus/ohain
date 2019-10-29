package org.monan.ohain.proxy;

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

    public Class get(Class source){
        ClassPool cp = ClassPool.getDefault() ;

        String className = source.getName()+"$Proxy" ;
        CtClass ctClass = cp.makeClass(className) ;

        Method[] methods = source.getDeclaredMethods() ;
        for(Method method : methods){
//            method.getReturnType()
//            CtMethod ctMethod = CtMethod.make("")
        }
        return null ;
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
