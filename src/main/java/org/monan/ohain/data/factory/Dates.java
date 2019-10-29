package org.monan.ohain.data.factory;

import java.util.Date;

/**
 * @author Yuewen.Huang
 * @email Yuewen.Huang@geely.com
 * @date 2019/10/23
 */
public class Dates implements Data<Date> {

    private Date minDate = new Date() ;

    @Override
    public <T> T get() {
        return null;
    }

    @Override
    public <T> T getBeta() {
        return null;
    }

    @Override
    public <T> T getAlpha() {

        // 获取未来时间
        return null;
    }
}
