package assurance.entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;
import javax.validation.constraints.NotNull;
import java.util.UUID;

/**
 * @author fdse
 */
@Data
@Entity
@GenericGenerator(name = "jpa-uuid", strategy = "uuid")
@JsonIgnoreProperties(ignoreUnknown = true)
public class Assurance {

    @Id
    @GeneratedValue(generator = "jpa-uuid")
    @Column(length = 32)
    private String id;

    /**
     * which order the assurance is related to
     */
    @NotNull
    private String orderId;

    /**
     * the type of assurance
     */
    @Enumerated(EnumType.STRING)
    private AssuranceType type;

    public Assurance(){
        this.orderId = UUID.randomUUID().toString();
    }

    public Assurance(String id, String orderId, AssuranceType type){
        this.id = id;
        this.orderId = orderId;
        this.type = type;
    }

}
