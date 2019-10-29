package org.monan.ohain.data.factory;

/**
 * @author Yuewen.Huang
 * @email Yuewen.Huang@geely.com
 * @date 2019/10/29
 */
public interface Data<T> {

    /**
     * 合格产品
     * @param <T>
     * @return
     */
    public <T> T get() ;


    /**
     * 边界
     * @param <T>
     * @return
     */
    public <T> T getBeta();


    /**
     * 不合格
     * @param <T>
     * @return
     */
    public <T> T getAlpha();
}
