package org.monan.ohain.data.factory;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.lang.reflect.Field;
import java.lang.reflect.Method;

/**
 * @author Yuewen.Huang
 * @email Yuewen.Huang@geely.com
 * @date 2019/10/23
 */
public class DataFactory{


    private static final Logger LOGGER = LoggerFactory.getLogger(DataFactory.class) ;

    public static DataFactory instance = new DataFactory() ;

    public static DataFactory getInstance(){
        return instance ;
    }

    public <T> T get(Class<T> clazz) {
        try{
            T instance = clazz.getDeclaredConstructor().newInstance() ;
            Field[] fields = clazz.getDeclaredFields() ;
            for(Field field : fields){
                Class type = field.getType() ;
                LOGGER.debug("field type {}",type.getName());
                //init
                Object value = null ;

                if (type.getName().equals("java.lang.String")){
                    value = new Strings().get() ;
                }else{
//                    value = manu(type);
                }

                // set property
                String fieldName = field.getName() ;
                LOGGER.debug("field name {}",fieldName);
                Method method = clazz.getMethod("set"+fieldName.substring(0,1).toUpperCase()+fieldName.substring(1),new Class[]{type}) ;
                if (method == null ){
                    //fixme
                } else {
                    method.invoke(instance,value) ;
                }

            }
            return instance ;
        } catch(Exception e) {
            throw new RuntimeException(e.getMessage(),e) ;
        }
    }

    public String getString(){
        return null ;
    }

    public Integer getInteger(){
        return null ;
    }


}
