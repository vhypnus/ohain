package org.monan.ohain.proxy;

import javassist.CannotCompileException;
import javassist.ClassPool;
import javassist.CtClass;
import javassist.CtMethod;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.lang.reflect.Method;
import java.lang.reflect.Parameter;

/**
 * @author Yuewen.Huang
 * @email Yuewen.Huang@geely.com
 * @date 2019/10/29
 */
public class Proxy {

    private static final Logger LOGGER = LoggerFactory.getLogger(Proxy.class) ;

    public Class get(Class source){
        ClassPool cp = ClassPool.getDefault() ;

        String className = source.getName()+"$Proxy" ;
        CtClass ctClass = cp.makeClass(className) ;

        Method[] methods = source.getDeclaredMethods() ;
        for(Method method : methods){
            if ("public".equals(getModifiers(method.getModifiers()))){
                String returnType = method.getReturnType().getName() ;
                String methodName = method.getName() ;
                String param = getParams(method.getParameters()) ;

                String retureStatment = "" ;
                if (!"void".equals(returnType)){
                    retureStatment = "return new Object() ;" ;
                }
                String format = "%s %s %s(%s){System.out.println(\"hello world\") ;%s}" ;
                String src = String.format(format,new String[]{"public",returnType,methodName,param,retureStatment}) ;
                LOGGER.debug(src);
                try {
                    CtMethod ctMethod = CtMethod.make(src,ctClass) ;
                    ctClass.addMethod(ctMethod);
                } catch (CannotCompileException e) {
                    LOGGER.error(e.getMessage(),e);
                    throw new RuntimeException(e.getMessage()) ;
                }


            }

        }

        Class clazz = null ;
        try {
            clazz = ctClass.toClass() ;
        } catch (CannotCompileException e) {
            LOGGER.error(e.getMessage(),e);
            throw new RuntimeException(e.getMessage()) ;
        }
        return clazz ;
    }


    private String getExceptions(){
        return null ;
    }

    private String getParams(Parameter[] parameters){
        StringBuilder params = new StringBuilder("") ;
        if (parameters != null && parameters.length > 0){
            for (int index = 0 ,max = parameters.length ; index < max ;index ++) {
                params.append(parameters[index].getClass().getName()).append(" " ).append("$"+index).append(",");
            }

            LOGGER.debug(params.toString());
            params.deleteCharAt(params.length()-1) ;
        }

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
