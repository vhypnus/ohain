package org.monan.ohain.data;

import java.util.Date;

/**
 * @author Yuewen.Huang
 * @email Yuewen.Huang@geely.com
 * @date 2019/10/23
 */
public class DateData implements Data<Date> {

    @Override
    public Date generate() {
        return new Date();
    }
}
